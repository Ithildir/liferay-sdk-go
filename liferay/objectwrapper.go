package liferay

const (
	orderByComparator string = "com.liferay.portal.kernel.util.OrderByComparator"
	serviceContext    string = "com.liferay.portal.service.ServiceContext"
)

type ObjectWrapper struct {
	ClassName string
	Object    map[string]interface{}
}

func (w ObjectWrapper) addClassName(params map[string]interface{}, name string, className string) {
	if len(w.ClassName) > 0 {
		className = w.ClassName
	}

	params[("+" + name)] = className
}

func (w ObjectWrapper) addObject(params map[string]interface{}, name string) {
	for k, v := range w.Object {
		params[(name + "." + k)] = v
	}
}

func (w ObjectWrapper) mangle(params map[string]interface{}, name string, className string) {
	w.addClassName(params, name, className)
	w.addObject(params, name)
}

func MangleObjectWrapper(params map[string]interface{}, name string, className string, wrapper *ObjectWrapper) {
	if wrapper == nil {
		if className != serviceContext {
			params[name] = nil
		}

		return
	}

	wrapper.mangle(params, name, className)
}
