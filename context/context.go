package context

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var Context = NewContext()

type MsContext struct {
	Request  *http.Request
	W        http.ResponseWriter
	routers  map[string]func(ctx *MsContext)
	pathArgs map[string]map[string]string
}

func NewContext() *MsContext {
	ctx := &MsContext{}
	ctx.routers = make(map[string]func(ctx2 *MsContext))
	ctx.pathArgs = make(map[string]map[string]string)
	return ctx
}

var UrlTree = NewTrie()

// 前缀树结构 用于路径参数匹配
type Trie struct {
	next   map[string]*Trie
	isWord bool
}

func NewTrie() Trie {
	root := new(Trie)
	root.next = make(map[string]*Trie)
	root.isWord = false
	return *root
}

// 插入数据， 路由根据 "/" 进行拆分
func (t *Trie) Insert(word string) {
	for _, v := range strings.Split(word, "/") {
		if t.next[v] == nil {
			node := new(Trie)
			node.next = make(map[string]*Trie)
			node.isWord = false
			t.next[v] = node
		}
		// * 匹配所有
		// {X}  匹配路由参数 X
		if v == "*" || strings.Index(v, "{") != -1 {
			t.isWord = true
		}
		t = t.next[v]
	}
	t.isWord = true
}

// 匹配路由
func (t *Trie) Search(word string) (isHave bool, arg map[string]string) {
	arg = make(map[string]string)
	isHave = false
	pathSegments := strings.Split(word, "/")

	// 遍历路径的每个部分
	for _, v := range pathSegments {
		if t.isWord {
			for k := range t.next {
				// 处理参数匹配，例如 `{pid}.html`
				if strings.HasPrefix(k, "{") {
					// 先去掉 `{}` 提取参数名称
					paramName := k[1:] // 去掉 `{`
					if strings.Contains(paramName, "}") {
						paramName = paramName[:strings.Index(paramName, "}")] // 去掉 `}`
					}

					// 如果 `k` 后面还有 `.html`，要特殊处理
					suffix := ""
					if strings.HasSuffix(k, ".html") {
						suffix = ".html"
					}

					// `v` 需要匹配 `.html` 结尾
					if suffix != "" {
						if !strings.HasSuffix(v, suffix) {
							return false, nil // 需要 `.html` 结尾
						}
						v = strings.TrimSuffix(v, suffix) // 去掉 `.html`
					}

					// 确保 `v` 作为参数值，不包含 `/`
					if strings.Contains(v, "/") {
						return false, nil
					}

					arg[paramName] = v // 存储参数
				} else if k != v {
					return false, nil // 确保路径完全匹配
				}
				v = k
			}
		}

		// 路径未匹配
		if t.next[v] == nil {
			return false, nil
		}
		t = t.next[v]
	}

	// 确保完全匹配整个路径
	if len(t.next) == 0 {
		isHave = t.isWord
	}
	return
}

// ServeHTTP 方法 - 处理 HTTP 请求
func (ctx *MsContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx.W = w
	ctx.Request = r
	path := r.URL.Path
	f := ctx.routers[path]

	if f == nil {
		for key, value := range ctx.routers {
			// 仅处理带 `{}` 的动态路径
			paramRouteRegex := regexp.MustCompile(`^(/[^{}]+)*(/{[^/{}]+})(/[^{}]*)*(\.html)?$`)
			if !paramRouteRegex.MatchString(key) {
				continue
			}

			// `.html` 结尾的路径，必须匹配 `.html`
			htmlRouteRegex := regexp.MustCompile(`.*\.html$`)
			if htmlRouteRegex.MatchString(path) != htmlRouteRegex.MatchString(key) {
				continue
			}

			// 进行路径参数匹配
			isHav, args := UrlTree.Search(path)
			if isHav {
				ctx.pathArgs[path] = args
				value(ctx)
				return
			}
		}
	} else {
		f(ctx)
	}
}

func (ctx *MsContext) Handler(url string, f func(context *MsContext)) {
	UrlTree.Insert(url)
	ctx.routers[url] = f
}

func (ctx *MsContext) GetPathVariable(key string) string {
	return ctx.pathArgs[ctx.Request.URL.Path][key]
}

func (ctx *MsContext) GetForm(key string) (string, error) {
	if err := ctx.Request.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		return "", err
	}
	return ctx.Request.Form.Get(key), nil
}
func (ctx *MsContext) GetJson(key string) interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	_ = json.Unmarshal(body, &params)
	return params[key]
}
