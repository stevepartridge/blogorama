package whitelist

var (
	list = map[string]struct{}{}
)

func AddMethod(method string) {
	method = clean(method)
	if exists(method) {
		return
	}

	list[method] = struct{}{}
}

func MethodAllowed(method string) bool {
	return exists(method)
}

func exists(method string) bool {
	method = clean(method)

	if _, ok := list[method]; ok {
		return true
	}
	return false
}

func clean(method string) string {
	if method != "" {
		if method[:1] == "/" {
			return method[1:]
		}
	}
	return method
}
