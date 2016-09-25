package tmpl_test

import (
	"github.com/bbuck/dragon-mud/scripting"
	. "github.com/bbuck/dragon-mud/text/tmpl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var script = `
    local tmpl = require("tmpl")

    tmpl.Register("Hello, %{name}!", "lua_test")

    function testTemplate()
        result, ok = tmpl.Render("lua_test", {name = "World"})

		return result, ok
    end
`

var _ = Describe("ScriptedTemplates", func() {
	var (
		e      *scripting.LuaEngine
		result string
		values []*scripting.LuaValue
		ok     bool
		err    error
	)

	e = scripting.NewLuaEngine()
	RegisterScriptModules(e)
	e.LoadString(script)

	BeforeEach(func() {
		values, err = e.Call("testTemplate", 2)
		if err == nil {
			result = values[0].AsString()
			ok = values[1].AsBool()
		}
	})

	It("successfully calls the script method", func() {
		Ω(err).Should(BeNil())
	})

	It("doesn't fail", func() {
		Ω(err).Should(BeNil())
		Ω(ok).Should(BeTrue())
	})

	It("should render correctly", func() {
		Ω(err).Should(BeNil())
		Ω(result).Should(Equal("Hello, World!"))
	})
})