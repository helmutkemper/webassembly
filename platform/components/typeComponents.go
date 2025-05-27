package components

import (
	"errors"
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/browser/stage"
	"github.com/helmutkemper/webassembly/mathUtil"
	"github.com/helmutkemper/webassembly/qrcode"
	"image/color"
	"log"
	"reflect"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

// todo: padrão da fábrica, autocomplete off
// todo: number de range, onFocusOut se value == null, valur = min

type searchStructRet struct {
	sliceElement   reflect.Value
	componentFound bool
	fieldFound     bool
	tagData        *tag
}

type Components struct {
	ref                interface{}
	panelFather        *html.TagDiv
	panelFatherContent *html.TagDiv
	panelBody          *html.TagDiv

	setupPanel       map[string]string
	setupBody        map[string]string
	setupPanelHeader map[string]string

	autoZIndex bool
	isDragging bool
	offsetX    int
	offsetY    int
	minimized  bool

	afterInit []func()
}

func (e *Components) CssPanelHeader(key, value string) {
	if e.setupPanelHeader == nil {
		e.setupPanelHeader = make(map[string]string)
	}

	e.setupPanelHeader[key] = value
}

func (e *Components) CssPanel(key, value string) {
	if e.setupPanel == nil {
		e.setupPanel = make(map[string]string)
	}

	e.setupPanel[key] = value
}

func (e *Components) CssBody(key, value string) {
	if e.setupBody == nil {
		e.setupBody = make(map[string]string)
	}

	e.setupBody[key] = value
}

func (e *Components) setupPanelHeaderInit() {
	if e.setupPanelHeader == nil {
		e.setupPanelHeader = make(map[string]string)
	}

	if _, found := e.setupPanelHeader["display"]; !found {
		e.setupPanelHeader["display"] = "flex"
	}

	if _, found := e.setupPanelHeader["alignItems"]; !found {
		e.setupPanelHeader["alignItems"] = "center"
	}

	//if _, found := e.setupPanelHeader["borderBottom"]; !found {
	//	e.setupPanelHeader["borderBottom"] = "1px solid #ccc"
	//}

	if _, found := e.setupPanelHeader["paddingBottom"]; !found {
		e.setupPanelHeader["paddingBottom"] = "5px"
	}
}

func (e *Components) setupBodyInit() {
	if e.setupBody == nil {
		e.setupBody = make(map[string]string)
	}

	if _, found := e.setupBody["display"]; !found {
		e.setupBody["display"] = "block"
	}

	if _, found := e.setupBody["maxHeight"]; !found {
		e.setupBody["maxHeight"] = "90vh"
	}

	if _, found := e.setupBody["flex"]; !found {
		e.setupBody["flex"] = "1 1 auto"
	}

	if _, found := e.setupBody["overflowY"]; !found {
		e.setupBody["overflowY"] = "auto"
	}

	if _, found := e.setupBody["padding"]; !found {
		e.setupBody["padding"] = "0 10px"
	}

	if _, found := e.setupBody["boxSizing"]; !found {
		e.setupBody["boxSizing"] = "border-box"
	}

	if _, found := e.setupBody["border"]; !found {
		e.setupBody["border"] = "none"
	}

	if _, found := e.setupBody["scrollbarWidth"]; !found {
		e.setupBody["scrollbarWidth"] = "thin"
	}

	if _, found := e.setupBody["scrollbarColor"]; !found {
		e.setupBody["scrollbarColor"] = "#888 transparent"
	}

	//if _, found := e.setupBody[""]; !found {
	//	e.setupBody[""] = ""
	//}

	e.applyCss()
}

func (e *Components) setupPanelInit() {
	if e.setupPanel == nil {
		e.setupPanel = make(map[string]string)
	}

	if _, found := e.setupPanel["textFontSize"]; !found {
		e.setupPanel["textFontSize"] = "12px"
	}

	if _, found := e.setupPanel["fontFamily"]; !found {
		e.setupPanel["fontFamily"] = "Arial, sans-serif"
	}

	if _, found := e.setupPanel["width"]; !found {
		e.setupPanel["width"] = "350px"
	}

	if _, found := e.setupPanel["borderRadius"]; !found {
		e.setupPanel["borderRadius"] = "10px"
	}

	if _, found := e.setupPanel["boxShadow"]; !found {
		e.setupPanel["boxShadow"] = "0 3px 6px rgba(0, 0, 0, 0.3)"
	}

	if _, found := e.setupPanel["border"]; !found {
		e.setupPanel["border"] = "1px solid #ccc"
	}

	if _, found := e.setupPanel["padding"]; !found {
		e.setupPanel["padding"] = "10px"
	}

	if _, found := e.setupPanel["backgroundColor"]; !found {
		e.setupPanel["backgroundColor"] = "#fff"
	}

	if _, found := e.setupPanel["maxHeight"]; !found {
		e.setupPanel["maxHeight"] = "80vh"
	}

	if _, found := e.setupPanel["flex"]; !found {
		e.setupPanel["flex"] = "0 0 auto"
	}

	if _, found := e.setupPanel["overflowY"]; !found {
		e.setupPanel["overflowY"] = "none"
	}

	if _, found := e.setupPanel["position"]; !found {
		e.setupPanel["position"] = "fixed"
	}

	if _, found := e.setupPanel["margin"]; !found {
		e.setupPanel["margin"] = "10px"
	}

	if _, found := e.setupPanel["flexDirection"]; !found {
		e.setupPanel["flexDirection"] = "column"
	}

	if _, found := e.setupPanel["display"]; !found {
		e.setupPanel["display"] = "flex"
	}

	//if _, found := e.setupPanel[""]; !found {
	//	e.setupPanel[""] = ""
	//}

	e.applyCss()
}

func (e *Components) applyCss() {
	if e.panelFather == nil {
		return
	}

	if e.panelBody == nil {
		return
	}

	for k := range e.setupPanel {
		e.panelFather.AddStyle(k, e.setupPanel[k])
	}

	for k := range e.setupBody {
		e.panelBody.AddStyle(k, e.setupBody[k])
	}

	head := new(html.TagHead).
		Css(".panel .compCel .inputNumber::-webkit-outer-spin-button, "+
			".panel .compCel .inputNumber::-webkit-inner-spin-button", "-webkit-appearance", "none").
		Css(".panel .compCel .inputNumber::-webkit-outer-spin-button, "+
			".panel .compCel .inputNumber::-webkit-inner-spin-button", "margin", 0).
		Css(".panel .compCel .inputNumber", "-moz-appearance", "textfield").
		Css(".panelBody::-webkit-scrollbar-thumb:hover", "background-color", "#555").
		Css(".panelBody::-webkit-scrollbar-thumb", "background-color", "#888").        // Bar color
		Css(".panelBody::-webkit-scrollbar-thumb", "border-radius", "10px").           // Rounded edges
		Css(".panelBody::-webkit-scrollbar-thumb", "border", "2px solid transparent"). // Spacing
		Css(".panelBody::-webkit-scrollbar-thumb", "background-clip", "content-box").  // Adjusts the background
		Css(".panelBody::-webkit-scrollbar-track", "background", "transparent").
		Css(".panelBody::-webkit-scrollbar", "width", "8px")
	head.CssAppend()
}

func (e *Components) addAfterInit(f func()) {
	e.afterInit = append(e.afterInit, f)
}

func (e *Components) runAfterInit() {
	for _, f := range e.afterInit {
		f()
	}
}

func (e *Components) Init(el any) (panel *html.TagDiv, err error) {
	e.afterInit = make([]func(), 0)

	element := reflect.ValueOf(el)
	typeof := reflect.TypeOf(el)
	e.setupPanelInit()
	e.setupBodyInit()
	e.setupPanelHeaderInit()

	e.createDivsFather()

	e.panelFather.HideForFade()

	err = e.process(element, typeof)
	if err != nil {
		//file, line, funcName := runTimeUtil.Trace()
		//err = errors.Join(fmt.Errorf("%v(line: %v).process().error: %v", funcName, line, err))
		//err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	e.runAfterInit()

	e.panelFather.Fade(200 * time.Millisecond)

	return e.panelFather, err
}

func (e *Components) createDivsFather() {
	e.autoZIndex = true

	e.panelFather = factoryBrowser.NewTagDiv().
		Class("panel").
		Data(map[string]string{"cell": "panel"})
	e.panelFather.AddStyle("userSelect", "none")
	e.panelFather.Get().Call("addEventListener", "mouseover", js.FuncOf(func(this js.Value, args []js.Value) any {
		if !e.autoZIndex {
			return nil
		}
		e.panelFather.AddStyle("zIndex", stage.GetNextZIndex())
		return nil
	}))
	e.panelFatherContent = factoryBrowser.NewTagDiv()
	e.panelFather.Append(e.panelFatherContent)

	e.panelBody = factoryBrowser.NewTagDiv().Class("panelBody")
	e.panelBody.AddStyle("display", "block")

	e.applyCss()
}

func (e *Components) process(element reflect.Value, typeof reflect.Type) (err error) {

	elementPrt := element

	if element.Kind() == reflect.Pointer {
		if element.CanSet() && element.IsNil() {
			newInstance := reflect.New(element.Type().Elem())
			element.Set(newInstance)
		}

		element = element.Elem()
		typeof = typeof.Elem()
	}

	if element.CanInterface() && element.Kind() == reflect.Struct {
		for i := 0; i != element.NumField(); i += 1 {
			fieldVal := element.Field(i)
			fieldTyp := typeof.Field(i)

			//if fieldVal.Kind() == reflect.Struct {
			//	e.process(fieldVal, fieldTyp.Type)
			//}

			tagRaw := fieldTyp.Tag.Get("wasmPanel")
			if tagRaw != "" {
				tagData := new(tag)
				if err = tagData.init(tagRaw); err != nil {
					err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
					return
				}

				switch tagData.Type {
				case "headerText":
					dragIcon, minimizeIcon, closeIcon := e.processHeaderText(fieldVal, tagData.Label, e.panelFatherContent)
					e.processHeaderTextAddDragListener(dragIcon)
					e.processHeaderTextAddMinimizeListener(minimizeIcon)
					e.processHeaderTextAddCloseListener(closeIcon)
					e.Show()
					// Espera criar panelHeader para que panelBody fique abaixo
					e.panelFather.Append(e.panelBody)
				case "panel":

					if tagData.Top == "" {
						tagData.Top = "5px"
					}

					if tagData.Left == "" {
						tagData.Left = "5px"
					}

					e.panelFather.AddStyle("top", tagData.Top)
					e.panelFather.AddStyle("left", tagData.Left)
					e.panelFather.AddStyle("zIndex", stage.GetNextZIndex())

					if fieldVal.Kind() == reflect.Pointer {
						if fieldVal.CanSet() && fieldVal.IsNil() {
							newInstance := reflect.New(fieldVal.Type().Elem())
							fieldVal.Set(newInstance)
						}

						e.ref = fieldVal.Interface()
					}

					if err = e.process(fieldVal, fieldTyp.Type); err != nil {
						return
					}

				case "compCel":
					// ignore
				case "component":
					divCompCel := factoryBrowser.NewTagDiv().
						Class("compCel").
						Data(map[string]string{"cell": "main"})

					err = e.processComponent(element, fieldVal, fieldTyp.Type, tagData, divCompCel)
					if err != nil {
						//file, line, funcName := runTimeUtil.Trace()
						//err = errors.Join(fmt.Errorf("%v(line: %v).processComponent().error: %v", funcName, line, err))
						//err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}

					panelCel := factoryBrowser.NewTagDiv().
						Class("panelCel").
						AddStyle("borderRadius", "10px").
						AddStyle("boxShadow", "0 3px 6px rgba(0, 0, 0, 0.3)").
						AddStyle("border", "1px solid #aaa").
						AddStyle("margin", "10px 0").
						AddStyle("padding", "10px").
						AddStyle("backgroundColor", "#f9f9f9").
						Data(map[string]string{"cell": "cell"})
					closeIcon := e.processLabelCel(tagData.Label, panelCel)
					e.processLabelCelAddCloseListener(closeIcon)

					panelCel.Append(
						divCompCel,
					)

					e.panelBody.Append(panelCel)
				case "celLabel":
					//log.Printf("%v", tagData.Type)
				case "celText":
					//log.Printf("%v", tagData.Type)
				case "range":
					//log.Printf("%v", tagData.Type)
				case "onchange":
					//log.Printf("%v", tagData.Type)
				default:
					//log.Printf(">>>>>>>> %v", tagData.Type)
					err = e.process(fieldVal, fieldTyp.Type)
					if err != nil {
						//file, line, funcName := runTimeUtil.Trace()
						//err = errors.Join(fmt.Errorf("%v(line: %v).process().error: %v", funcName, line, err))
						//err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}
				}

				switch tagData.Type {
				case "panel", "component":
					// Put the context menu on the main panel
					if err = e.processContextMenu(elementPrt, element, fieldVal, fieldTyp, e.panelFatherContent); err != nil {
						return
					}

				case "mainMenu":
					if err = e.processMainMenu(elementPrt, element, fieldVal, fieldTyp); err != nil {
						return
					}
				}
			}
		}
	}

	return
}

func (e *Components) GetFather() (father *html.TagDiv) {
	return e.panelFather
}

func (e *Components) processContextMenu(parentElementPtr, parentElement, element reflect.Value, typeof reflect.StructField, father html.Compatible) (err error) {

	// The menu is created for each element contained in the panel (concept failure)
	// This code checks and avoids multiple menus
	//menuElement := parentElement.FieldByName("ContextMenu")
	//if menuElement.IsValid() {
	//	menuMethod := menuElement.MethodByName("GetMenu")
	//	if menuMethod.IsValid() {
	//		ops := menuMethod.Call(nil)
	//		if !ops[0].IsZero() {
	//			return
	//		}
	//	}
	//}

	var componentMenuData searchStructRet
	componentMenuData = e.searchStruct(parentElementPtr, "ContextMenu", "Type", "contextMenu")
	a := componentMenuData.componentFound
	b := componentMenuData.fieldFound

	if !(a || b) {
		return
	}

	if !element.CanInterface() {
		err = fmt.Errorf("component.Menu (%v.%+v) cannot be transformed into an interface", parentElement.Type().Name(), typeof.Name)
		return
	}

	if element.Kind() != reflect.Pointer {
		err = fmt.Errorf("component.Menu (%v) requires a pointer to the component, example", parentElement.Type().Name())
		err = errors.Join(err, fmt.Errorf("     type %v struct {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("       components.ContextMenu"))
		err = errors.Join(err, fmt.Errorf("       "))
		err = errors.Join(err, fmt.Errorf("       %v *%v `wasmPanel:\"type:contextMenu;...\"`", typeof.Name, typeof.Type))
		err = errors.Join(err, fmt.Errorf("     }"))
		err = errors.Join(err, fmt.Errorf("     func (e *%v) InitMenu() {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("       e.MenuData = getMenuData()"))
		err = errors.Join(err, fmt.Errorf("     }"))
		return
	}

	// Checks if the import of `components.ContextMenu` is fine, using XOR logic
	if (a || b) && !(a && b) {
		err = fmt.Errorf("error: component %v needs to embed `components.ContextMenu` directly", parentElement.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.ContextMenu"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         MenuData *[]MenuOptions `wasmPanel:\"type:contextMenu;func:InitMenu\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *%v) InitMenu() {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         e.MenuData = getMenuData()"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	}

	if componentMenuData.tagData.Func == "" {
		err = fmt.Errorf("error: component %v has an embed `components.ContextMenu` correctily", parentElement.Type().Name())
		err = errors.Join(err, fmt.Errorf("       However, `wasmPanel:\"type:contextMenu\"` does not contain the name of the function to configure the menu when this initializate"))
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.ContextMenu"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         MenuData *[]MenuOptions `wasmPanel:\"type:contextMenu;func:InitMenu\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *%v) InitMenu() {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         e.MenuData = getMenuData()"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	}

	menuMethod := parentElementPtr.MethodByName(componentMenuData.tagData.Func)
	if !menuMethod.IsValid() {
		err = fmt.Errorf("error: component %v has an embed `components.ContextMenu` correctily", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       However, `wasmPanel:\"type:contextMenu\"` contain a function that cannot be called."))
		err = errors.Join(err, fmt.Errorf("       See if the name contained in the tag is the same name as the function."))
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.ContextMenu"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         MenuData *[]MenuOptions `wasmPanel:\"type:contextMenu;func:%v\"`", componentMenuData.tagData.Func))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *%v) InitMenu() {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         e.MenuData = getMenuData()"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	}

	var menuOptions []options
	parentElementPtr.MethodByName(componentMenuData.tagData.Func).Call(nil)
	if menuOptions, err = e.processSliceMenuOptions(element, componentMenuData.sliceElement); err != nil {
		return
	}

	if componentMenuData.tagData.Columns == "" {
		componentMenuData.tagData.Columns = "2"
	}

	node := father.Get()

	f := func() {
		node = e.dataSetNavigate(node)

		// comment block - start
		// The menu is created for each element within the panel (initial project concept failure)
		// This code makes only one contextual menu added to node
		// Concept failure: If a second contextual menu is added, it will be ignored
		if cell := node.Get("dataset").Get("hasContextMenu"); !cell.IsUndefined() && cell.String() != "" {
			return
		}

		node.Get("dataset").Set("hasContextMenu", "true")
		// comment block - end

		contextMenu := ContextMenu{}
		contextMenu.Menu(menuOptions)
		contextMenu.AttachMenuJs(node)
		contextMenu.Css("gridGridTemplateColumns", fmt.Sprintf("repeat(%v, 1fr)", componentMenuData.tagData.Columns))
		contextMenu.Init()

		menuElement := parentElement.FieldByName("ContextMenu")
		if !menuElement.CanSet() {
			err = fmt.Errorf("error: component %v has an embed `components.ContextMenu` correctily", element.Type().Name())
			err = errors.Join(err, fmt.Errorf("         However, there was a problem when transferring the menu created to the existing component"))
			return
		}

		menuElement.Set(reflect.ValueOf(contextMenu))
	}
	e.addAfterInit(f)

	return
}

// dataSetNavigate Navigates to Div Father, Cell or Panel
//
// This function can only be used after the appendChild() command, because before it the div is floating, so it is
// performed inside an afterInit() function block
func (e *Components) dataSetNavigate(node js.Value) (append js.Value) {
	if cell := node.Get("dataset").Get("cell"); !cell.IsUndefined() {
		if cell.String() == "panel" || cell.String() == "cell" {
			return node
		}
	}

	parent := node.Get("parentElement")
	if !parent.IsNull() {
		return e.dataSetNavigate(parent)
	}

	parent = node.Get("parentNode")
	if !parent.IsNull() {
		return e.dataSetNavigate(parent)
	}

	return node
}

func (e *Components) processMainMenu(parentElementPtr, parentElement, element reflect.Value, typeof reflect.StructField) (err error) {

	var componentMenuData searchStructRet
	componentMenuData = e.searchStruct(parentElementPtr, "MainMenu", "Type", "mainMenu")
	a := componentMenuData.componentFound
	b := componentMenuData.fieldFound

	if !(a || b) {
		return
	}

	if !element.CanInterface() {
		err = fmt.Errorf("component.Menu (%v.%+v) cannot be transformed into an interface", parentElement.Type().Name(), typeof.Name)
		return
	}

	if element.Kind() != reflect.Pointer {
		err = fmt.Errorf("component.Menu (%v) requires a pointer to the component, example", parentElement.Type().Name())
		err = errors.Join(err, fmt.Errorf("     type %v struct {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("       components.MainMenu"))
		err = errors.Join(err, fmt.Errorf("       "))
		err = errors.Join(err, fmt.Errorf("       %v *%v `wasmPanel:\"type:mainMenu;...\"`", typeof.Name, typeof.Type))
		err = errors.Join(err, fmt.Errorf("     }"))
		err = errors.Join(err, fmt.Errorf("     func (e *%v) InitMenu() {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("       e.MenuData = getMenuData()"))
		err = errors.Join(err, fmt.Errorf("     }"))
		return
	}

	// Checks if the import of `components.ContextMenu` is fine, using XOR logic
	if (a || b) && !(a && b) {
		err = fmt.Errorf("error: component %v needs to embed `components.ContextMenu` directly", parentElement.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.MainMenu"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         MenuData *[]MenuOptions `wasmPanel:\"type:mainMenu;func:InitMenu;left:10px;top:10px\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *%v) InitMenu() {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         e.MenuData = getMenuData()"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	}

	if componentMenuData.tagData.Func == "" {
		err = fmt.Errorf("error: component %v has an embed `components.ContextMenu` correctily", parentElement.Type().Name())
		err = errors.Join(err, fmt.Errorf("       However, `wasmPanel:\"type:mainMenu\"` does not contain the name of the function to configure the menu when this initializate"))
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.MainMenu"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         MenuData *[]MenuOptions `wasmPanel:\"type:mainMenu;func:InitMenu;left:10px;top:10px\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *%v) InitMenu() {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         e.MenuData = getMenuData()"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	}

	menuMethod := parentElementPtr.MethodByName(componentMenuData.tagData.Func)
	if !menuMethod.IsValid() {
		err = fmt.Errorf("error: component %v has an embed `components.ContextMenu` correctily", parentElement.Type().Name())
		err = errors.Join(err, fmt.Errorf("       However, `wasmPanel:\"type:mainMenu\"` contain a function that cannot be called."))
		err = errors.Join(err, fmt.Errorf("       See if the name contained in the tag is the same name as the function."))
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.MainMenu"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         MenuData *[]MenuOptions `wasmPanel:\"type:mainMenu;func:%v;left:10px;top:10px\"`", componentMenuData.tagData.Func))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *%v) InitMenu() {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         e.MenuData = getMenuData()"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	}

	var menuOptions []options
	parentElementPtr.MethodByName(componentMenuData.tagData.Func).Call(nil)
	if menuOptions, err = e.processSliceMenuOptions(element, componentMenuData.sliceElement); err != nil {
		return
	}

	if componentMenuData.tagData.Columns == "" {
		componentMenuData.tagData.Columns = "2"
	}

	mainMenu := MainMenu{}
	mainMenu.Menu(menuOptions)
	mainMenu.FixedMenu(componentMenuData.tagData.Left, componentMenuData.tagData.Top)
	mainMenu.Css("menuTitle", componentMenuData.tagData.Label)
	mainMenu.Css("gridGridTemplateColumns", fmt.Sprintf("repeat(%v, 1fr)", componentMenuData.tagData.Columns))
	mainMenu.Init()

	menuElement := parentElement.FieldByName("MainMenu")
	if !menuElement.CanSet() {
		err = fmt.Errorf("error: component %v has an embed `components.ContextMenu` correctily", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("         However, there was a problem when transferring the menu created to the existing component"))
		return
	}

	menuElement.Set(reflect.ValueOf(mainMenu))
	return
}

func (e *Components) debugStruct(element reflect.Value) {

	log.SetPrefix("debug struct: >>>>>>>>>>\n")
	defer log.SetPrefix("")

	log.Printf("Debugging struct or pointer to struct")
	if element.Kind() == reflect.Ptr {
		log.Printf("element is a pointer")
		if element.IsNil() {
			log.Printf("pointer is nil")
			return
		}
		element = element.Elem()
	}

	if element.Kind() == reflect.Struct {
		log.Printf("element is a struct: %s", element.Type().Name())
		for i := 0; i < element.NumField(); i++ {
			field := element.Field(i)
			fieldType := element.Type().Field(i)
			log.Printf("Field Name: %s,\nField Type: %s,\nField Value: %+v", fieldType.Name, field.Type(), field.Interface())
		}
	} else {
		log.Printf("element is not a struct or a pointer to struct, kind: %s", element.Kind())
	}
}

// searchStruct
//
// English:
//
//	Search for the struct of the injected component in the struct father, and a struct that has the tag determined by
//	the KET/VALUE key. E.g. Component: "Menu" and struct key "Type", struct value on the key "contextMenu"
//
// Português:
//
//	Procura pelo struct do componente injetado no struct pai e por um struct que tenha a tag determinada pela
//	chave ket/value. Ex.: E.g. Componente: "Menu" e struct chave "Type", struct valor da chave "contextMenu"
func (e *Components) searchStruct(searchElement reflect.Value, componentName, tagKey, tagValue string) (objectData searchStructRet) {

	objectData.tagData = new(tag)

	if searchElement.Kind() == reflect.Ptr {
		if searchElement.IsNil() {
			return
		}
		searchElement = searchElement.Elem()
	}

	if searchElement.Kind() == reflect.Struct {
		for i := 0; i < searchElement.NumField(); i++ {
			fieldType := searchElement.Type().Field(i)
			objectData.sliceElement = searchElement.Field(i)

			if fieldType.Name == componentName {
				objectData.componentFound = true
			}

			tagRaw := fieldType.Tag.Get("wasmPanel")
			if tagRaw != "" {
				_ = objectData.tagData.init(tagRaw)
				objectData.fieldFound = objectData.tagData.Search(tagKey, tagValue)
				if objectData.fieldFound && objectData.componentFound {
					return
				}
			}
		}
	}

	return
}

func (e *Components) processSliceMenuOptions(element, sliceData reflect.Value) (menuOptions []options, err error) {
	A := sliceData.Kind() == reflect.Ptr && sliceData.Elem().Kind() == reflect.Slice
	B := sliceData.Kind() == reflect.Slice
	if !(A || B) {
		return
	}

	tagData := new(tag)
	menuOptions = make([]options, 0)

	if sliceData.Kind() == reflect.Ptr {
		sliceData = sliceData.Elem()
	}

	if sliceData.Kind() == reflect.Slice {
		for i := 0; i < sliceData.Len(); i++ {
			item := sliceData.Index(i)

			menuOptItem := options{}

			for j := 0; j < item.NumField(); j++ {
				field := item.Field(j)
				fieldType := item.Type().Field(j)

				tagRaw := fieldType.Tag.Get("wasmPanel")
				if tagRaw != "" {
					if err = tagData.init(tagRaw); err != nil {
						err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
						return
					}

					switch tagData.Type {
					case "label":
						menuOptItem.Label = field.String()
					case "icon":
						menuOptItem.Icon = field.String()
					case "iconLeft":
						menuOptItem.IconLeft = field.String()
					case "iconRight":
						menuOptItem.IconRight = field.String()
					case "type":
						menuOptItem.Type = field.String()
					case "options":
						if menuOptItem.Items, err = e.processSliceMenuOptions(element, field); err != nil {
							return
						}
					case "action":
						menuOptItem.Action = field.Interface().(js.Func)
					case "subMenu":
						if menuOptItem.Submenu, err = e.processSliceMenuOptions(element, field); err != nil {
							return
						}
					}
				}
			}

			menuOptions = append(menuOptions, menuOptItem)
		}
	}
	return
}

func (e *Components) processComponent(parentElement, element reflect.Value, typeof reflect.Type, tagData *tag, father *html.TagDiv) (err error) {

	boardComponent := Board{}
	elementPrt := element

	if !element.CanInterface() {
		err = fmt.Errorf("component.Board (%v) cannot be transformed into an interface", parentElement.Type().Name())
		return
	}

	if element.Kind() != reflect.Pointer {
		err = fmt.Errorf("component.Board (%v) requires a pointer to the component, example", parentElement.Type().Name())
		err = errors.Join(err, fmt.Errorf("type %v struct {", parentElement.Type().Name()))
		err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:range;label:...\"`", typeof.Name(), element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("}"))
		return
	}

	// populates the pointer, if it is nil
	if element.CanSet() && element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// passes from pointer to element
	//elementPrt := element
	element = element.Elem()

	// Checks if the import of `components.Board` was done
	if fieldBoard := element.FieldByName("Board"); !fieldBoard.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Board` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Board"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Dragging *DraggingEffect   `wasmPanel:\"type:range;label:effect\"`"))
		err = errors.Join(err, fmt.Errorf("         Tween    *TweenSelect      `wasmPanel:\"type:select;label:Tween function\"`"))
		err = errors.Join(err, fmt.Errorf("         Start    *EasingTweenStart `wasmPanel:\"type:button;label:start easing tween\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Board
		newInstance := reflect.New(fieldBoard.Type())
		fieldBoard.Set(newInstance.Elem())

		// Initializes the two input tags within Board
		boardComponent.__divTag = father

		// populates the component.Board within the user component
		componentBoard := element.FieldByName("Board")
		componentBoard.Set(reflect.ValueOf(boardComponent))
	}

	//if fieldBoard := element.FieldByName("Menu"); fieldBoard.IsValid() {
	//	log.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//}

	//contextMenuFound := false
	//for i := 0; i != element.NumField(); i += 1 {
	//	var fieldTyp reflect.StructField
	//	//fieldVal := element.Field(i)
	//	if typeof.Kind() == reflect.Pointer {
	//		fieldTyp = typeof.Elem().Field(i)
	//	} else {
	//		fieldTyp = typeof.Field(i)
	//	}
	//
	//	tagRaw := fieldTyp.Tag.Get("wasmPanel")
	//	if tagRaw != "" {
	//		tagData := new(tag)
	//		if err = tagData.init(tagRaw); err != nil {
	//			err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
	//			return
	//		}
	//
	//		switch tagData.Type {
	//		case "contextMenu":
	//			contextMenuFound = true
	//		}
	//	}
	//}

	for i := 0; i != element.NumField(); i += 1 {
		divComponent := factoryBrowser.NewTagDiv().Class("component").
			AddStyle("pointerEvents", "auto").
			AddStyle("display", "flex").
			AddStyle("align-items", "center").
			AddStyle("margin", "5px 0").
			Data(map[string]string{"cell": "component"})

		var fieldTyp reflect.StructField
		fieldVal := element.Field(i)
		if typeof.Kind() == reflect.Pointer {
			fieldTyp = typeof.Elem().Field(i)
		} else {
			fieldTyp = typeof.Field(i)
		}

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagData := new(tag)
			if err = tagData.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagData.Type {
			case "range":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Range (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:range;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Range (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentRange(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "osm":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Osm (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:osm;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Osm (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentOsm(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "button":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Button (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:button;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Button (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentButton(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "select":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Select (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:select;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Select (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentSelect(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "radio":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Radio (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:radio;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Radio (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentRadio(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "checkbox":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Checkbox (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:checkbox;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Checkbox (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentCheckbox(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "color":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Color (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:color;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Color (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentColor(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "date":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Date (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:date;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Date (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentDate(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "week":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Week (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:week;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Week (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentWeek(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "text":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Text (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:text;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Text (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentText(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "qrcode":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.QRCode (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:qrcode;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.QRCode (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentQRCode(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "url":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Url (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:url;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Url (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentUrl(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "tel":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Tel (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:tel;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Tel (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentTel(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "email":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Email (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:email;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Email (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentMail(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "time":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Time (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:time;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Time (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentTime(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "month":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Month (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:month;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Month (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentMonth(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "password":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Password (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:password;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Password (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentPassword(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "mail":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.Mail (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:mail;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.Mail (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentMail(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			case "textArea":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("component.TextArea (%v) requires a pointer to the component, example", fieldVal.Type().Name())
					err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
					err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:text;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
					err = errors.Join(err, fmt.Errorf("}"))
					return
				}

				if !fieldVal.CanSet() || !fieldVal.CanInterface() {
					err = fmt.Errorf("component.TextArea (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
					return
				}

				err = e.processComponentTextArea(fieldVal, tagData, divComponent)
				if err != nil {
					return
				}

			default:
				//log.Printf(">>>>>>>> %v", tagData.Type)
			}

			switch tagData.Type {
			case "range", "osm", "button", "select", "radio", "checkbox", "color",
				"date", "week", "text", "qrcode", "url", "tel", "email", "time",
				"month", "password", "mail", "textArea":
				if err = e.processContextMenu(elementPrt, element, fieldVal, fieldTyp, divComponent); err != nil {
					return
				}

			default:
				//log.Printf(">>>>>>>> %v", tagData.Type)
			}
		}

		father.Append(divComponent)
	}

	return
}

func (e *Components) processLabelCelAddCloseListener(closeIcon *html.TagDiv) {
	closeIcon.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		args[0].Call("stopPropagation")

		node := this.Get("parentNode").Get("parentNode")
		children := node.Get("children")
		child := children.Index(1)

		display := child.Get("style").Get("display").String()
		if display != "none" {
			child.Get("dataset").Set("display", display)
			child.Get("style").Set("display", "none")
			return nil
		}

		display = child.Get("dataset").Get("display").String()
		child.Get("style").Set("display", display)
		child.Get("dataset").Set("display", "")
		return nil
	}))
}

func (e *Components) processLabelCel(label string, father *html.TagDiv) (closeIcon *html.TagDiv) {
	// <div class="labelCel">
	//   <div class="labelText">Label</div>
	//   <div class="closeIcon">ˇ</div>
	// </div>

	closeIcon = factoryBrowser.NewTagDiv().Class("closeIcon").Text("ˇ")
	father.Append(
		factoryBrowser.NewTagDiv().Class("labelCel").
			AddStyle("display", "flex").
			AddStyle("justify-content", "space-between").
			AddStyle("align-items", "center").
			AddStyle("border-bottom", "1px solid #ddd").
			AddStyle("padding-bottom", 0).
			AddStyle("margin-bottom", 0).
			Append(
				factoryBrowser.NewTagDiv().Class("labelText").
					AddStyle("font-weight", "bold").Text(label),
				closeIcon,
			),
	)

	return
}

//func (e *Components) callFunc(funcObj reflect.Value, params ...interface{}) (results []interface{}, err error) {
//
//	if funcObj.Kind() != reflect.Func {
//		return nil, fmt.Errorf("funcObj is not of type reflect.Func")
//	}
//
//	if len(params) != funcObj.Type().NumIn() {
//		return nil, fmt.Errorf("incorrect number of parameters")
//	}
//
//	in := make([]reflect.Value, len(params))
//	for i, param := range params {
//		in[i] = reflect.ValueOf(param)
//	}
//
//	out := funcObj.Call(in)
//
//	results = make([]interface{}, len(out))
//	for i, result := range out {
//		results[i] = result.Interface()
//	}
//	return results, nil
//}

// searchFieldByTagType Procura um campo pelo tipo da tag
func (e *Components) searchFieldByTagType(typeName, eventName string, element reflect.Value) (fieldVal reflect.Value, fieldTyp reflect.StructField, found bool) {
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal = element.Field(i)
		fieldTyp = reflect.TypeOf(element.Interface()).Field(i)

		tagDataInternal := new(tag)
		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			if err := tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				log.Printf("%v", err) // todo: melhorar isto
				return
			}
		}

		if tagDataInternal.Type == typeName && eventName == "" {
			found = true
			return
		}

		if tagDataInternal.Type == typeName && tagDataInternal.Event == eventName {
			found = true
			return
		}
	}

	return
}

func (e *Components) verifyTypeFromElement(fieldVal reflect.Value, valueType reflect.Kind) (dataType reflect.Kind, value any, ok bool) {
	switch valueType {
	case reflect.Int64:
		dataType = reflect.Int64
		value = fieldVal.Int()
	case reflect.Float64:
		dataType = reflect.Float64
		value = fieldVal.Float()
	case reflect.String:
		dataType = reflect.String
		value = fieldVal.String()
	case reflect.Bool:
		dataType = reflect.Bool
		value = fieldVal.Bool()
	default:
		dataType = reflect.Invalid
		return
	}

	ok = true
	return
}

func (e *Components) verifyTypeFloat64Element(fieldVal reflect.Value, valueType reflect.Kind) (dataType reflect.Kind, value any, ok bool) {
	switch valueType {
	case reflect.Float64:
		dataType = reflect.Float64
		value = fieldVal.Float()
	default:
		dataType = reflect.Invalid
		return
	}

	ok = true
	return
}

func (e *Components) verifyTypeInt64Element(fieldVal reflect.Value, valueType reflect.Kind) (dataType reflect.Kind, value any, ok bool) {
	switch valueType {
	case reflect.Int64:
		dataType = reflect.Int64
		value = fieldVal.Int()
	default:
		dataType = reflect.Invalid
		return
	}

	ok = true
	return
}

//func (e *Components) verifyTypeNumericFromElement(fieldVal reflect.Value, valueType reflect.Kind) (dataType reflect.Kind, value any, ok bool) {
//	switch valueType {
//	case reflect.Int64:
//		dataType = reflect.Int64
//		value = fieldVal.Int()
//	case reflect.Float64:
//		dataType = reflect.Float64
//		value = fieldVal.Float()
//	default:
//		dataType = reflect.Invalid
//		return
//	}
//
//	ok = true
//	return
//}

func (e *Components) processComponentOsm(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	_ = dataType
	var value any
	var ok bool

	var longitude, latitude float64
	var zoom int
	var url string
	_ = url

	elementOriginal := element
	osmComponent := Osm{}

	//tagCanvas := new(html.TagCanvas)
	//ref.Init(width, height)
	//tagCanvas.Id(mathUtil.GetUID())

	// todo: mudar nome de inputOsm
	inputOsm := factoryBrowser.NewTagDiv().Class("inputOsm").
		Data(map[string]string{"cell": "osm"}) //.Append(tagCanvas)

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Osm` was done
	if fieldOsm := element.FieldByName("Osm"); !fieldOsm.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Osm` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Osm"))
		err = errors.Join(err, fmt.Errorf("         "))
		// todo: colocar texto correto
		err = errors.Join(err, fmt.Errorf("         Value int64 `wasmPanel:\"type:value;min:0;max:50;step:1;default:0\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Osm
		newInstance := reflect.New(fieldOsm.Type())
		fieldOsm.Set(newInstance.Elem())

		// Initializes the two input tags within Osm
		osmComponent.__osmTag = inputOsm
		osmComponent.__canvasTag = new(html.TagCanvas)

		// __osmOnInputEvent is the pointer sent when the `change` event happens
		osmComponent.__change = new(__osmOnInputEvent)

		// populates the component.Osm within the user component
		componentOsm := element.FieldByName("Osm")
		componentOsm.Set(reflect.ValueOf(osmComponent))
	}

	// Search for the listener input tag and if it does not exist, set up the controller control function
	//if _, _, found := e.searchFieldByTagType("listener", "input", element); !found {
	//	var methods []reflect.Value
	//	var params []interface{}
	//	log.Printf("osmComponent.__canvasTag: %v", osmComponent.__canvasTag)
	//	// Passes the functions to be executed in the listener
	//	methods = []reflect.Value{
	//		// osmComponent is the struct components.Osm and OnChangeNumber is a function belonging to the struct components.Osm
	//		//todo: fazer eventos
	//		reflect.ValueOf(osmComponent.__change).MethodByName("OnMousedown"),
	//	}
	//
	//	// Pass variable pointers
	//	params = []interface{}{
	//		// __osmOnInputEvent is the type pointer contained in components.Osm and collects value
	//		osmComponent.__canvasTag,
	//	}
	//
	//	// explanation
	//	//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
	//	inputOsm.ListenerAddReflect("mousedown", params, methods, element.Interface())
	//}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input osm tag was requested by the user
			case "tagOsm":
				fieldVal.Set(reflect.ValueOf(inputOsm))

			// Checks whether the reference to the canvas tag was requested by the user
			case "tagCanvas":
				//osmComponent.__canvasTag = new(html.TagCanvas)
				fieldVal.Set(reflect.ValueOf(osmComponent.__canvasTag))

			// Checks if the value tag was created
			case "latitude":

				// Captures the value of the component defined by the latitude tag
				dataType, value, ok = e.verifyTypeFloat64Element(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type float64", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Fill in the numeric fields
				latitude = value.(float64)

				// If the value is zero, and the user has determined a value other than zero,
				// fill in the field with the default value
				if latitude == 0 && tagDataInternal.Default != "" {

					var defaultValue float64
					defaultValue, err = strconv.ParseFloat(tagDataInternal.Default, 64)
					if err != nil {
						err = fmt.Errorf("%v.%v type '%v', must be a default value type string of float. Found: %v", element.Type().Name(), fieldTyp.Name, fieldVal.Kind(), tagDataInternal.Default)
						return
					}

					latitude = defaultValue
				}

				//log.Printf("latitude: %v", latitude)

			// Checks if the value tag was created
			case "longitude":

				// Captures the value of the component defined by the latitude tag
				dataType, value, ok = e.verifyTypeFloat64Element(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type float64", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Fill in the numeric fields
				longitude = value.(float64)

				// If the value is zero, and the user has determined a value other than zero,
				// fill in the field with the default value
				if longitude == 0 && tagDataInternal.Default != "" {

					var defaultValue float64
					defaultValue, err = strconv.ParseFloat(tagDataInternal.Default, 64)
					if err != nil {
						err = fmt.Errorf("%v.%v type '%v', must be a default value type string of float. Found: %v", element.Type().Name(), fieldTyp.Name, fieldVal.Kind(), tagDataInternal.Default)
						return
					}

					longitude = defaultValue
				}

				//log.Printf("longitude: %v", longitude)

			case "zoom":

				// Captures the value of the component defined by the latitude tag
				dataType, value, ok = e.verifyTypeInt64Element(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int or int64", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Fill in the numeric fields
				zoom = int(value.(int64))

				// If the value is zero, and the user has determined a value other than zero,
				// fill in the field with the default value
				if zoom == 0 && tagDataInternal.Default != "" {

					var defaultValue int64
					defaultValue, err = strconv.ParseInt(tagDataInternal.Default, 10, 64)
					if err != nil {
						//todo: mensagem de erro
						return
					}

					zoom = int(defaultValue)

					if tagDataInternal.Width == "" || tagDataInternal.Height == "" {
						// todo: colocar o erro
						return
					}

					var width, height int64
					width, err = strconv.ParseInt(tagDataInternal.Width, 10, 64)
					if err != nil {
						// todo: colocar o erro
						return
					}

					height, err = strconv.ParseInt(tagDataInternal.Height, 10, 64)
					if err != nil {
						// todo: colocar o erro
						return
					}

					//if tagCanvas == nil {
					//	tagCanvas = new(html.TagCanvas)
					//}

					//log.Printf("width: %v", width)
					//log.Printf("height: %v", height)
					osmComponent.__canvasTag.Init(int(width), int(height))
					osmComponent.__canvasTag.Id(mathUtil.GetUID())

				}

				log.Printf("zoom: %v", zoom)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				_ = methods
				var params []interface{}
				_ = params
				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				// explanation
				//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
				//   The first element of the array is the user function
				//   From the second element onwards, they are internal functions and must be called after the user function in case the user has changed any value.
				//inputRange.ListenerAddReflect(tagDataInternal.Event, params, methods, element.Interface())
				//inputNumber.ListenerAddReflect(tagDataInternal.Event, params, methods, element.Interface())

				//inputRange.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
				//inputNumber.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)

			}
		}
	}

	if longitude != 0 && latitude != 0 && zoom != 0 {
		longitude = -48.465279 + 0.000400000
		latitude = -27.428942
		zoom = 18
	}

	//if tagCanvas == nil {
	//	tagCanvas = new(html.TagCanvas)
	//	tagCanvas.Init(int(250), int(250))
	//	tagCanvas.Id(mathUtil.GetUID())
	//}

	inputOsm.Append(osmComponent.__canvasTag)
	osmComponent.__canvasTag.SetOsm(longitude, latitude, zoom, 0, 0)

	osmComponent.__canvasTag.Get().Call("addEventListener", "mousedown", js.FuncOf(osmComponent.onMouseDown))

	father.Append(
		//factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputOsm,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Osm{}) {
			r := fieldVal.Interface().(Osm)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentRange(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	rangeComponent := Range{}

	inputRange := factoryBrowser.NewTagInputRange().Class("inputRange").
		AddStyle("flex", 2).
		AddStyle("margin", "0 10px") // todo: remover daqui
	inputNumber := factoryBrowser.NewTagInputNumber().Class("inputNumber").
		AddStyle("width", "80px").
		AddStyle("textAlign", "center") // todo: remover daqui

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Range` was done
	if fieldRange := element.FieldByName("Range"); !fieldRange.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Range` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Range"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value int64 `wasmPanel:\"type:value;min:0;max:50;step:1;default:0\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Range
		newInstance := reflect.New(fieldRange.Type())
		fieldRange.Set(newInstance.Elem())

		// Initializes the two input tags within Range
		rangeComponent.__rangeTag = inputRange
		rangeComponent.__numberTag = inputNumber

		// __rangeOnInputEvent is the pointer sent when the `change` event happens
		rangeComponent.__change = new(__rangeOnInputEvent)

		// populates the component.Range within the user component
		componentRange := element.FieldByName("Range")
		componentRange.Set(reflect.ValueOf(rangeComponent))
	}

	// Search for the listener input tag and if it does not exist, set up the controller control function
	if _, _, found := e.searchFieldByTagType("listener", "input", element); !found {
		var methods []reflect.Value
		var params []interface{}

		// Passes the functions to be executed in the listener
		methods = []reflect.Value{
			// rangeComponent is the struct components.Range and OnChangeNumber is a function belonging to the struct components.Range
			reflect.ValueOf(&rangeComponent).MethodByName("OnChangeNumber"),
			// rangeComponent is the struct components.Range and OnChangeRange is a function belonging to the struct components.Range
			reflect.ValueOf(&rangeComponent).MethodByName("OnChangeRange"),
		}

		// Pass variable pointers
		params = []interface{}{
			// __rangeOnInputEvent is the type pointer contained in components.Range and collects value
			new(__rangeOnInputEvent),
			// __rangeOnInputEvent is the type pointer contained in components.Range and collects value
			new(__rangeOnInputEvent),
		}

		// explanation
		//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
		inputRange.ListenerAddReflect("input", params, methods, element.Interface())
		inputNumber.ListenerAddReflect("input", params, methods, element.Interface())
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input range tag was requested by the user
			case "inputTagRange":
				fieldVal.Set(reflect.ValueOf(inputRange))
				//fieldVal.Interface().(*html.TagInputRange).Min(-5).Max(10).Value(-5)

			// Checks whether the reference to the input number tag was requested by the user
			case "inputTagNumber":
				fieldVal.Set(reflect.ValueOf(inputNumber))
				//fieldVal.Interface().(*html.TagInputNumber).Min(-5).Max(10).Value(-5)

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.Float64:
					if value.(float64) != 0 {
						passValue = true
					}

				case reflect.Int64:
					if value.(int64) != 0 {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type int64 or float64", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Fill in the numeric fields
				inputRange.Value(value)
				inputNumber.Value(value)

				// If the value is zero, and the user has determined a value other than zero,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputRange.Value(tagDataInternal.Default)
					inputNumber.Value(tagDataInternal.Default)
				}

				// todo: isto não tem sentido aqui - inicio
				switch dataType {
				case reflect.Float64:
					if tagDataInternal.Min != "" {
						eMin, _ := strconv.ParseFloat(tagDataInternal.Min, 64)
						inputRange.Min(eMin)
						inputNumber.Min(eMin)
					}
					if tagDataInternal.Max != "" {
						eMax, _ := strconv.ParseFloat(tagDataInternal.Max, 64)
						inputRange.Max(eMax)
						inputNumber.Max(eMax)
					}
					if tagDataInternal.Step != "" {
						eStep, _ := strconv.ParseFloat(tagDataInternal.Step, 64)
						inputRange.Step(eStep)
						inputNumber.Step(eStep)
					}
				case reflect.Int64:
					if tagDataInternal.Min != "" {
						eMin, _ := strconv.ParseInt(tagDataInternal.Min, 10, 64)
						inputRange.Min(eMin)
						inputNumber.Min(eMin)
					}
					if tagDataInternal.Max != "" {
						eMax, _ := strconv.ParseInt(tagDataInternal.Max, 10, 64)
						inputRange.Max(eMax)
						inputNumber.Max(eMax)
					}
					if tagDataInternal.Step != "" {
						eStep, _ := strconv.ParseInt(tagDataInternal.Step, 10, 64)
						inputRange.Step(eStep)
						inputNumber.Step(eStep)
					}
				default:

				}
				// todo: isto não tem sentido aqui - fim

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				// explanation
				//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
				//   The first element of the array is the user function
				//   From the second element onwards, they are internal functions and must be called after the user function in case the user has changed any value.
				//inputRange.ListenerAddReflect(tagDataInternal.Event, params, methods, element.Interface())
				//inputNumber.ListenerAddReflect(tagDataInternal.Event, params, methods, element.Interface())

				inputRange.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
				inputNumber.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)

			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputRange,
		inputNumber,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Range{}) {
			r := fieldVal.Interface().(Range)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentColor(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	colorComponent := Color{}

	inputColor := factoryBrowser.NewTagInputColor().Class("component .component-color")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Color` was done
	if fieldColor := element.FieldByName("Color"); !fieldColor.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Color` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Color"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed color\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Color
		newInstance := reflect.New(fieldColor.Type())
		fieldColor.Set(newInstance.Elem())

		// Initializes the two input tags within Color
		colorComponent.__colorTag = inputColor

		// __colorOnInputEvent is the pointer sent when the `change` event happens
		colorComponent.__change = new(__colorOnInputEvent)

		// populates the component.Color within the user component
		componentColor := element.FieldByName("Color")
		componentColor.Set(reflect.ValueOf(colorComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input color tag was requested by the user
			case "inputTagColor":
				fieldVal.Set(reflect.ValueOf(inputColor))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputColor.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputColor.Value(tagDataInternal.Default)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputColor.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputColor,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Color{}) {
			r := fieldVal.Interface().(Color)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentDate(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	dateComponent := Date{}

	inputDate := factoryBrowser.NewTagInputDate().Class("component .component-date")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Date` was done
	if fieldDate := element.FieldByName("Date"); !fieldDate.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Date` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Date"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed date;placeHolder:Place holder date\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Date
		newInstance := reflect.New(fieldDate.Type())
		fieldDate.Set(newInstance.Elem())

		// Initializes the two input tags within Date
		dateComponent.__dateTag = inputDate

		// __dateOnInputEvent is the pointer sent when the `change` event happens
		dateComponent.__change = new(__dateOnInputEvent)

		// populates the component.Date within the user component
		componentDate := element.FieldByName("Date")
		componentDate.Set(reflect.ValueOf(dateComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input date tag was requested by the user
			case "inputTagDate":
				fieldVal.Set(reflect.ValueOf(inputDate))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputDate.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputDate.Value(tagDataInternal.Default)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputDate.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputDate,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Date{}) {
			r := fieldVal.Interface().(Date)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentText(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	textComponent := Text{}

	inputText := factoryBrowser.NewTagInputText().Class("component component-text")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Text` was done
	if fieldText := element.FieldByName("Text"); !fieldText.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Text` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Text"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed text;placeHolder:Place holder text\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Text
		newInstance := reflect.New(fieldText.Type())
		fieldText.Set(newInstance.Elem())

		// Initializes the two input tags within Text
		textComponent.__textTag = inputText

		// __textOnInputEvent is the pointer sent when the `change` event happens
		textComponent.__change = new(__textOnInputEvent)

		// populates the component.Text within the user component
		componentText := element.FieldByName("Text")
		componentText.Set(reflect.ValueOf(textComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input text tag was requested by the user
			case "inputTagText":
				fieldVal.Set(reflect.ValueOf(inputText))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputText.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputText.Value(tagDataInternal.Default)
				}

				inputText.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputText.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputText,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Text{}) {
			r := fieldVal.Interface().(Text)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentQRCode(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var fieldComponent reflect.Value
	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	qrCodeComponent := QRCode{}

	//tagCanvas := factoryBrowser.NewTagCanvas(255,255)//.Class("component .component-text")
	tagCanvas := new(html.TagCanvas)

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Text` was done
	if fieldText := element.FieldByName("QRCode"); !fieldText.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Text` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.QRCode"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;size:512;default:'htts://www.google.com'\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize QRCode
		fieldComponent = fieldText
	}

	var qrDisableBorder bool
	var qrCodeSize int
	var qrCodeRecoveryLevel int
	var qrCodeColor color.Color
	var qrCodeBackground color.Color
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			case "disableBorder":

				if !fieldVal.IsValid() {

				} else if disableBorder, ok := fieldVal.Interface().(bool); ok {
					qrDisableBorder = disableBorder

				} else {
					err = fmt.Errorf("%v.%v type '%v', must be a type bool", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

			case "color":
				if !fieldVal.IsValid() {

				} else if color, ok := fieldVal.Interface().(string); ok {
					if color == "" {
						color = "#000000"
					}

					qrCodeColor, err = mathUtil.HexToColor(color)
					if err != nil {
						err = fmt.Errorf("%v.%v type '%v', contains an error in the value: %v", element.Type().Name(), fieldTyp.Name, fieldVal.Kind(), err)
						return
					}
				} else {
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

			case "background":
				if !fieldVal.IsValid() {

				} else if color, ok := fieldVal.Interface().(string); ok {
					if color == "" {
						color = "#ffffff"
					}

					qrCodeBackground, err = mathUtil.HexToColor(color)
					if err != nil {
						err = fmt.Errorf("%v.%v type '%v', contains an error in the value: %v", element.Type().Name(), fieldTyp.Name, fieldVal.Kind(), err)
						return
					}
				} else {
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

			case "size":
				if !fieldVal.IsValid() {

				} else if size, ok := fieldVal.Interface().(int); ok {
					qrCodeSize = size
				} else {
					err = fmt.Errorf("%v.%v type '%v', must be a type int", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

			case "level":

				if !fieldVal.IsValid() || fieldVal.Int() == 0 {
					qrCodeRecoveryLevel = 2
				} else if level, ok := fieldVal.Interface().(int); ok {

					if level < 1 || level > 4 {
						err = fmt.Errorf("%v.%v type '%v(%v)'", element.Type().Name(), fieldTyp.Name, fieldVal.Kind(), fieldVal.Interface())
						err = errors.Join(err, fmt.Errorf("  values:"))
						err = errors.Join(err, fmt.Errorf("    1 - Level Low: 7%% error recovery"))
						err = errors.Join(err, fmt.Errorf("    2 - Level Medium: 15%% error recovery. Good default choice"))
						err = errors.Join(err, fmt.Errorf("    3 - Level High: 25%% error recovery"))
						err = errors.Join(err, fmt.Errorf("    4 - Level Highest: 30%% error recovery."))
						return
					}
					qrCodeRecoveryLevel = level
				} else {
					err = fmt.Errorf("%v.%v type '%v', must be a type int", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					err = errors.Join(err, fmt.Errorf("  values:"))
					err = errors.Join(err, fmt.Errorf("    1 - Level Low: 7%% error recovery"))
					err = errors.Join(err, fmt.Errorf("    2 - Level Medium: 15%% error recovery. Good default choice"))
					err = errors.Join(err, fmt.Errorf("    3 - Level High: 25%% error recovery"))
					err = errors.Join(err, fmt.Errorf("    4 - Level Highest: 30%% error recovery."))
					return
				}
			}
		}
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the canvas tag was requested by the user
			case "tagCanvas":
				fieldVal.Set(reflect.ValueOf(tagCanvas))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				switch dataType {
				case reflect.String:

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				//if tagDataInternal.Size == "" {
				//	err = fmt.Errorf("%v.%v tag config 'size' must be set with numeric value, eg. 512", element.Type().Name(), fieldTyp.Name)
				//	return
				//}

				if qrCodeSize == 0 && tagDataInternal.Size != "" {
					var size int64
					size, err = strconv.ParseInt(tagDataInternal.Size, 10, 64)
					if err != nil {
						err = fmt.Errorf("%v.%v tag config 'size' return an error: %v", element.Type().Name(), fieldTyp.Name, err)
						return
					}
					qrCodeSize = int(size)
				}

				if qrCodeSize == 0 {
					qrCodeSize = 298
				}
				tagCanvas.Init(qrCodeSize, qrCodeSize)

				if qrCodeRecoveryLevel == 0 && tagDataInternal.Level != "" {
					var level int64
					level, err = strconv.ParseInt(tagDataInternal.Level, 10, 64)
					if err != nil {
						err = fmt.Errorf("%v.%v tag config 'level' return an error: %v", element.Type().Name(), fieldTyp.Name, err)
						return
					}

					if level < 1 || level > 4 {
						err = fmt.Errorf("%v.%v config 'level', error", element.Type().Name(), fieldTyp.Name)
						err = errors.Join(err, fmt.Errorf("  values:"))
						err = errors.Join(err, fmt.Errorf("    1 - Level Low: 7%% error recovery"))
						err = errors.Join(err, fmt.Errorf("    2 - Level Medium: 15%% error recovery. Good default choice"))
						err = errors.Join(err, fmt.Errorf("    3 - Level High: 25%% error recovery"))
						err = errors.Join(err, fmt.Errorf("    4 - Level Highest: 30%% error recovery."))
						return
					}

					// 0 is used to know if level was defined, therefore, level must be greater than zero
					qrCodeRecoveryLevel = int(level) - 1
				} else if qrCodeRecoveryLevel == 0 {
					qrCodeRecoveryLevel = 1
				} else {
					// level was defined, but 0 is used to know if level was defined, therefore, level must be greater than zero
					qrCodeRecoveryLevel -= 1
				}

				if qrCodeColor == nil && tagDataInternal.Color != "" {
					qrCodeColor, err = mathUtil.HexToColor(tagDataInternal.Color)
					if err != nil {
						err = fmt.Errorf("%v.%v config 'color', error: %v", element.Type().Name(), fieldTyp.Name, err)
						return
					}
				} else if qrCodeColor == nil {
					qrCodeColor = color.Black
				}

				if qrCodeBackground == nil && tagDataInternal.Background != "" {
					qrCodeBackground, err = mathUtil.HexToColor(tagDataInternal.Background)
					if err != nil {
						err = fmt.Errorf("%v.%v config 'color', error: %v", element.Type().Name(), fieldTyp.Name, err)
						return
					}
				} else if qrCodeBackground == nil {
					qrCodeBackground = color.White
				}

				qrCodeComponent.__size = qrCodeSize
				qrCodeComponent.__recoveryLevel = qrcode.RecoveryLevel(qrCodeRecoveryLevel)
				qrCodeComponent.__background = qrCodeBackground
				qrCodeComponent.__color = qrCodeColor
				qrCodeComponent.__disableBorder = qrDisableBorder

				if converted, ok := value.(string); ok && converted != "" {
					tagCanvas.DrawQRCodeColor(qrCodeSize, converted, qrcode.RecoveryLevel(qrCodeRecoveryLevel), qrCodeColor, qrCodeBackground, qrDisableBorder)
				} else {
					tagCanvas.DrawQRCodeColor(qrCodeSize, tagDataInternal.Default, qrcode.RecoveryLevel(qrCodeRecoveryLevel), qrCodeColor, qrCodeBackground, qrDisableBorder)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				tagCanvas.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		//factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		tagCanvas,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Text{}) {
			r := fieldVal.Interface().(Text)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	// ------------------------------------------------------------------------------------------------------------------

	// Initialize QRCode
	newInstance := reflect.New(fieldComponent.Type())
	fieldComponent.Set(newInstance.Elem())

	// Initializes the canvas tag
	qrCodeComponent.__canvasTag = tagCanvas

	// __textOnInputEvent is the pointer sent when the `change` event happens
	qrCodeComponent.__change = new(__qrCodeOnInputEvent)

	// populates the component.Text within the user component
	componentQRCode := element.FieldByName("QRCode")
	componentQRCode.Set(reflect.ValueOf(qrCodeComponent))
	return
}

func (e *Components) processComponentUrl(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	urlComponent := Url{}

	inputUrl := factoryBrowser.NewTagInputUrl().Class("component .component-url")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Url` was done
	if fieldUrl := element.FieldByName("Url"); !fieldUrl.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Url` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Url"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed url;placeHolder:Place holder url\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Url
		newInstance := reflect.New(fieldUrl.Type())
		fieldUrl.Set(newInstance.Elem())

		// Initializes the two input tags within Url
		urlComponent.__urlTag = inputUrl

		// __urlOnInputEvent is the pointer sent when the `change` event happens
		urlComponent.__change = new(__urlOnInputEvent)

		// populates the component.Url within the user component
		componentUrl := element.FieldByName("Url")
		componentUrl.Set(reflect.ValueOf(urlComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input url tag was requested by the user
			case "inputTagUrl":
				fieldVal.Set(reflect.ValueOf(inputUrl))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputUrl.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputUrl.Value(tagDataInternal.Default)
				}

				inputUrl.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputUrl.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputUrl,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Url{}) {
			r := fieldVal.Interface().(Url)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentTel(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	telComponent := Tel{}

	inputTel := factoryBrowser.NewTagInputTel().Class("component .component-tel")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Tel` was done
	if fieldTel := element.FieldByName("Tel"); !fieldTel.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Tel` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Tel"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed tel;placeHolder:Place holder tel\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Tel
		newInstance := reflect.New(fieldTel.Type())
		fieldTel.Set(newInstance.Elem())

		// Initializes the two input tags within Tel
		telComponent.__telTag = inputTel

		// __telOnInputEvent is the pointer sent when the `change` event happens
		telComponent.__change = new(__telOnInputEvent)

		// populates the component.Tel within the user component
		componentTel := element.FieldByName("Tel")
		componentTel.Set(reflect.ValueOf(telComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input tel tag was requested by the user
			case "inputTagTel":
				fieldVal.Set(reflect.ValueOf(inputTel))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputTel.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputTel.Value(tagDataInternal.Default)
				}

				inputTel.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputTel.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputTel,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Tel{}) {
			r := fieldVal.Interface().(Tel)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentMail(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	mailComponent := Mail{}

	inputMail := factoryBrowser.NewTagInputEMail().
		Class("component-mail").
		AddStyle("display", "flex").
		AddStyle("justifyContent", "space-between").
		AddStyle("alignItems", "center").
		AddStyle("margin-left", "10px").
		AddStyle("width", "calc(60% - 20px)")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Mail` was done
	if fieldMail := element.FieldByName("Mail"); !fieldMail.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Mail` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Mail"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed mail;placeHolder:Place holder mail\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Mail
		newInstance := reflect.New(fieldMail.Type())
		fieldMail.Set(newInstance.Elem())

		// Initializes the two input tags within Mail
		mailComponent.__mailTag = inputMail

		// __mailOnInputEvent is the pointer sent when the `change` event happens
		mailComponent.__change = new(__mailOnInputEvent)

		// populates the component.Mail within the user component
		componentMail := element.FieldByName("Mail")
		componentMail.Set(reflect.ValueOf(mailComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input mail tag was requested by the user
			case "inputTagMail":
				fieldVal.Set(reflect.ValueOf(inputMail))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputMail.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputMail.Value(tagDataInternal.Default)
				}

				inputMail.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputMail.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputMail,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Mail{}) {
			r := fieldVal.Interface().(Mail)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentTime(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	timeComponent := Time{}

	inputTime := factoryBrowser.NewTagInputTime().Class("component .component-time")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Time` was done
	if fieldTime := element.FieldByName("Time"); !fieldTime.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Time` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Time"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed time;placeHolder:Place holder time\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Time
		newInstance := reflect.New(fieldTime.Type())
		fieldTime.Set(newInstance.Elem())

		// Initializes the two input tags within Time
		timeComponent.__timeTag = inputTime

		// __timeOnInputEvent is the pointer sent when the `change` event happens
		timeComponent.__change = new(__timeOnInputEvent)

		// populates the component.Time within the user component
		componentTime := element.FieldByName("Time")
		componentTime.Set(reflect.ValueOf(timeComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input time tag was requested by the user
			case "inputTagTime":
				fieldVal.Set(reflect.ValueOf(inputTime))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputTime.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					d, err := time.ParseDuration(tagDataInternal.Default)
					if err != nil {
						inputTime.Value(tagDataInternal.Default)
					} else {
						inputTime.Value(Timespan(d).Format(time.TimeOnly))
					}

				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputTime.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputTime,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Time{}) {
			r := fieldVal.Interface().(Time)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentMonth(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	monthComponent := Month{}

	inputMonth := factoryBrowser.NewTagInputMonth().Class("component .component-month")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Month` was done
	if fieldMonth := element.FieldByName("Month"); !fieldMonth.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Month` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Month"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed month;placeHolder:Place holder month\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Month
		newInstance := reflect.New(fieldMonth.Type())
		fieldMonth.Set(newInstance.Elem())

		// Initializes the two input tags within Month
		monthComponent.__monthTag = inputMonth

		// __monthOnInputEvent is the pointer sent when the `change` event happens
		monthComponent.__change = new(__monthOnInputEvent)

		// populates the component.Month within the user component
		componentMonth := element.FieldByName("Month")
		componentMonth.Set(reflect.ValueOf(monthComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input month tag was requested by the user
			case "inputTagMonth":
				fieldVal.Set(reflect.ValueOf(inputMonth))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputMonth.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputMonth.Value(tagDataInternal.Default)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputMonth.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputMonth,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Month{}) {
			r := fieldVal.Interface().(Month)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentWeek(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	weekComponent := Week{}

	inputWeek := factoryBrowser.NewTagInputWeek().Class("component .component-week")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Week` was done
	if fieldWeek := element.FieldByName("Week"); !fieldWeek.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Week` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Week"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed week;placeHolder:Place holder week\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Week
		newInstance := reflect.New(fieldWeek.Type())
		fieldWeek.Set(newInstance.Elem())

		// Initializes the two input tags within Week
		weekComponent.__weekTag = inputWeek

		// __weekOnInputEvent is the pointer sent when the `change` event happens
		weekComponent.__change = new(__weekOnInputEvent)

		// populates the component.Week within the user component
		componentWeek := element.FieldByName("Week")
		componentWeek.Set(reflect.ValueOf(weekComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input week tag was requested by the user
			case "inputTagWeek":
				fieldVal.Set(reflect.ValueOf(inputWeek))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputWeek.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputWeek.Value(tagDataInternal.Default)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputWeek.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputWeek,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Week{}) {
			r := fieldVal.Interface().(Week)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentPassword(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	passwordComponent := Password{}

	inputPassword := factoryBrowser.NewTagInputPassword().Class("component .component-password")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Password` was done
	if fieldPassword := element.FieldByName("Password"); !fieldPassword.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Password` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Password"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed password;placeHolder:Place holder text\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Password
		newInstance := reflect.New(fieldPassword.Type())
		fieldPassword.Set(newInstance.Elem())

		// Initializes the two input tags within Password
		passwordComponent.__passwordTag = inputPassword

		// __passwordOnInputEvent is the pointer sent when the `change` event happens
		passwordComponent.__change = new(__passwordOnInputEvent)

		// populates the component.Password within the user component
		componentPassword := element.FieldByName("Password")
		componentPassword.Set(reflect.ValueOf(passwordComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input password tag was requested by the user
			case "inputTagPassword":
				fieldVal.Set(reflect.ValueOf(inputPassword))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputPassword.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputPassword.Value(tagDataInternal.Default)
				}

				inputPassword.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputPassword.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputPassword,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Password{}) {
			r := fieldVal.Interface().(Password)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentTextArea(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	textAreaComponent := TextArea{}

	inputTextArea := factoryBrowser.NewTagTextArea().Class("component .component-textarea")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.TextArea` was done
	if fieldTextArea := element.FieldByName("TextArea"); !fieldTextArea.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.TextArea` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.TextArea"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed text;placeHolder:Place holder text\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize TextArea
		newInstance := reflect.New(fieldTextArea.Type())
		fieldTextArea.Set(newInstance.Elem())

		// Initializes the textArea tag within TextArea
		textAreaComponent.__textAreaTag = inputTextArea

		// __textAreaOnInputEvent is the pointer sent when the `change` event happens
		textAreaComponent.__change = new(__textAreaOnInputEvent)

		// populates the component.TextArea within the user component
		componentTextArea := element.FieldByName("TextArea")
		componentTextArea.Set(reflect.ValueOf(textAreaComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the textArea tag was requested by the user
			case "inputTagTextArea":
				fieldVal.Set(reflect.ValueOf(inputTextArea))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputTextArea.Text(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputTextArea.Text(tagDataInternal.Default)
				}

				inputTextArea.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputTextArea.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagDataFather.Label),
		inputTextArea,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(TextArea{}) {
			r := fieldVal.Interface().(TextArea)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentButton(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	buttonComponent := Button{}

	inputButton := factoryBrowser.NewTagInputButton().
		Class("inputButton")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Button` was done
	if fieldButton := element.FieldByName("Button"); !fieldButton.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Button` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type OnClickEvent struct {"))
		err = errors.Join(err, fmt.Errorf("         IsTrusted bool `wasmGet:\"isTrusted\"`"))
		err = errors.Join(err, fmt.Errorf("         Value     string `wasmGet:\"value\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *OnClickEvent, ref %v) OnClick(event OnClickEvent) {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         log.Printf(\"Trusted: %%v\", event.IsTrusted)"))
		err = errors.Join(err, fmt.Errorf("         log.Printf(\"Value:   %%v\", event.Value)"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Button"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Label    string        `wasmPanel:\"type:value;default:Ok\"`"))
		err = errors.Join(err, fmt.Errorf("         RunEvent *OnClickEvent `wasmPanel:\"type:listener;event:click;func:OnClick\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Range
		newInstance := reflect.New(fieldButton.Type())
		fieldButton.Set(newInstance.Elem())

		// Initializes the two input tags within Range
		buttonComponent.__buttonTag = inputButton

		// populates the component.Range within the user component
		componentButton := element.FieldByName("Button")
		componentButton.Set(reflect.ValueOf(buttonComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {

			// Checks whether the reference to the input range tag was requested by the user
			case "inputTagButton":
				fieldVal.Set(reflect.ValueOf(inputButton))
				//fieldVal.Interface().(*html.TagInputButton)

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type int64 or float64", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Fill in the numeric fields
				inputButton.Value(value)

				// If the value is zero, and the user has determined a value other than zero,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Label != "" {
					inputButton.Value(tagDataInternal.Label)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					// tagDataInternal.Func is the user function
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					// fieldVal.Interface() is the struct pointer that collects user data
					fieldVal.Interface(),
				}

				// explanation
				//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
				//   The first element of the array is the user function
				//   From the second element onwards, they are internal functions and must be called after the user function in case the user has changed any value.
				//inputButton.ListenerAddReflect(tagDataInternal.Event, params, methods, element.Interface())
				inputButton.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)

			}

		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagData.Label),
		inputButton,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Button{}) {
			r := fieldVal.Interface().(Button)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

// verifyTypesComponentSelect checks the types of the components.Select configuration
func (e *Components) verifyTypesComponentSelect(element reflect.Value) (err error) {

	elemTpl := element.Type()
	elemTplOriginal := element.Type()
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := elemTpl.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {
			case "value":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("the field %v, inside %v, must be a pointer of slice struct", fieldTyp.Name, elemTplOriginal.Name())
					err = errors.Join(err, fmt.Errorf("       Example:"))
					err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
					err = errors.Join(err, fmt.Errorf("         components.Select"))
					err = errors.Join(err, fmt.Errorf("         "))
					err = errors.Join(err, fmt.Errorf("         %v *[]SelectData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name))
					err = errors.Join(err, fmt.Errorf("       }"))
					err = errors.Join(err, fmt.Errorf("       "))
					err = errors.Join(err, fmt.Errorf("       type SelectData struct {"))
					err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
					err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
					err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
					err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
					err = errors.Join(err, fmt.Errorf("       }"))

					return
				}

				if fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.Elem().Kind() != reflect.Slice {
					err = fmt.Errorf("the field %v, inside %v, must be a pointer of slice struct", fieldTyp.Name, elemTplOriginal.Name())
					err = errors.Join(err, fmt.Errorf("       Example:"))
					err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
					err = errors.Join(err, fmt.Errorf("         components.Select"))
					err = errors.Join(err, fmt.Errorf("         "))
					err = errors.Join(err, fmt.Errorf("         %v *[]SelectData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name))
					err = errors.Join(err, fmt.Errorf("       }"))
					err = errors.Join(err, fmt.Errorf("       "))
					err = errors.Join(err, fmt.Errorf("       type SelectData struct {"))
					err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
					err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
					err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
					err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
					err = errors.Join(err, fmt.Errorf("       }"))

					return
				}

				fieldVal = fieldVal.Elem()
				fieldTpl := fieldVal.Type().Elem()

				if fieldTpl.Kind() != reflect.Struct {
					err = fmt.Errorf("the field %v, inside %v, must be a pointer of slice struct", fieldTyp.Name, elemTplOriginal.Name())
					err = errors.Join(err, fmt.Errorf("       Example:"))
					err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
					err = errors.Join(err, fmt.Errorf("         components.Select"))
					err = errors.Join(err, fmt.Errorf("         "))
					err = errors.Join(err, fmt.Errorf("         %v *[]SelectData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name))
					err = errors.Join(err, fmt.Errorf("       }"))
					err = errors.Join(err, fmt.Errorf("       "))
					err = errors.Join(err, fmt.Errorf("       type SelectData struct {"))
					err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
					err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
					err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
					err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
					err = errors.Join(err, fmt.Errorf("       }"))

					return
				}

				for k := 0; k != fieldTpl.NumField(); k += 1 {
					fieldTyp := fieldTpl.Field(k)

					tagRaw := fieldTyp.Tag.Get("wasmPanel")
					if tagRaw != "" {
						tagDataInternal := new(tag)
						if err = tagDataInternal.init(tagRaw); err != nil {
							err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
							return
						}

						switch tagDataInternal.Type {
						case "label":
							if fieldTyp.Type.Kind() != reflect.String {
								err = fmt.Errorf("the tag type:%v, inside %v, must be a string", tagDataInternal.Type, fieldTpl.Name())
								err = errors.Join(err, fmt.Errorf("       Example:"))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
								err = errors.Join(err, fmt.Errorf("         components.Select"))
								err = errors.Join(err, fmt.Errorf("         "))
								err = errors.Join(err, fmt.Errorf("         %v *[]%v `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name, fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("       }"))
								err = errors.Join(err, fmt.Errorf("       "))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
								err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
								err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
								err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
								err = errors.Join(err, fmt.Errorf("       }"))

								return
							}
						case "value":
							if fieldTyp.Type.Kind() != reflect.String {
								err = fmt.Errorf("the tag type:%v, inside %v, must be a string", tagDataInternal.Type, fieldTpl.Name())
								err = errors.Join(err, fmt.Errorf("       Example:"))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
								err = errors.Join(err, fmt.Errorf("         components.Select"))
								err = errors.Join(err, fmt.Errorf("         "))
								err = errors.Join(err, fmt.Errorf("         %v *[]%v `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name, fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("       }"))
								err = errors.Join(err, fmt.Errorf("       "))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
								err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
								err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
								err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
								err = errors.Join(err, fmt.Errorf("       }"))

								return
							}
						case "disabled":
							if fieldTyp.Type.Kind() != reflect.Bool {
								err = fmt.Errorf("the tag type:%v, inside %v, must be a boolean", tagDataInternal.Type, fieldTpl.Name())
								err = errors.Join(err, fmt.Errorf("       Example:"))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
								err = errors.Join(err, fmt.Errorf("         components.Select"))
								err = errors.Join(err, fmt.Errorf("         "))
								err = errors.Join(err, fmt.Errorf("         %v *[]%v `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name, fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("       }"))
								err = errors.Join(err, fmt.Errorf("       "))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
								err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
								err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
								err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
								err = errors.Join(err, fmt.Errorf("       }"))

								return
							}
						case "selected":
							if fieldTyp.Type.Kind() != reflect.Bool {
								err = fmt.Errorf("the tag type:%v, inside %v, must be a boolean", tagDataInternal.Type, fieldTpl.Name())
								err = errors.Join(err, fmt.Errorf("       Example:"))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
								err = errors.Join(err, fmt.Errorf("         components.Select"))
								err = errors.Join(err, fmt.Errorf("         "))
								err = errors.Join(err, fmt.Errorf("         %v *[]%v `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name, fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("       }"))
								err = errors.Join(err, fmt.Errorf("       "))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
								err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
								err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
								err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
								err = errors.Join(err, fmt.Errorf("       }"))

								return
							}
						}

					}
				}
				return
			}

		}

	}

	return
}

func (e *Components) processComponentSelect(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	inputSelect := factoryBrowser.NewTagSelect().Class("inputSelect")

	elementOriginal := element
	selectComponent := Select{}

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move the element from pointer to struct
	element = element.Elem()

	// Checks if the import of `components.Select` was done
	if fieldSelect := element.FieldByName("Select"); !fieldSelect.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Select` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Select"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         List *[]SelectData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label 3,value 3\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       type SelectData struct {"))
		err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
		err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
		err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       // Note: Use `>` to set value as selected. ie. >label,value"))
		return
	} else {
		// Initialize Select
		newInstance := reflect.New(fieldSelect.Type())
		fieldSelect.Set(newInstance.Elem())

		// Initializes the input tags within Select
		selectComponent.__selectTag = inputSelect

		// __selectOnInputEvent is the pointer sent when the `change` event happens
		selectComponent.__change = new(__selectOnInputEvent)

		// populates the component.Select within the user component
		componentRange := element.FieldByName("Select")
		componentRange.Set(reflect.ValueOf(selectComponent))
	}

	err = e.verifyTypesComponentSelect(element)
	if err != nil {
		return
	}

	fieldTyp := element.Type()
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {
			case "inputTagSelect":
				fieldVal.Set(reflect.ValueOf(inputSelect))

			case "value":

				// pointer is not nil
				// Move the element from pointer to struct
				fieldVal = fieldVal.Elem()

				if fieldVal.Len() == 0 {

					if tagDataInternal.Default != "" {
						optionList := strings.Split(tagDataInternal.Default, ",")
						if len(optionList)%2 != 0 {
							err = fmt.Errorf("%v.%v: the correct format from tag value is: `wasmPanel:\"type:value;default:label1,value1,label2,value2,labelN,valueN\"`, where value and label, must be a pair", element.Type().Name(), fieldTyp.Name)
							return
						}

						for k := 0; k != len(optionList); k += 2 {
							// if label start with `>` the option is selected
							selected := false
							if strings.HasPrefix(optionList[k], ">") {
								optionList[k] = optionList[k][1:]
								selected = true
							}
							inputSelect.NewOption(optionList[k], optionList[k+1], false, selected)
						}
					}

				} else {

					// run inside slice data
					for iField := 0; iField != fieldVal.Len(); iField += 1 {
						keyVal := fieldVal.Index(iField)

						// get label, value, disabled and selected
						var label, value string
						var disabled, selected bool
						for ik := 0; ik != keyVal.NumField(); ik += 1 {
							optionVal := keyVal.Field(ik)
							optionTyp := reflect.TypeOf(keyVal.Interface()).Field(ik)

							optionTagRaw := optionTyp.Tag.Get("wasmPanel")
							if optionTagRaw != "" {
								optionTag := new(tag)
								optionTag.init(optionTagRaw)

								switch optionTag.Type {
								case "label":
									label = optionVal.Interface().(string)
								case "value":
									value = optionVal.Interface().(string)
								case "disabled":
									disabled = optionVal.Interface().(bool)
								case "selected":
									selected = optionVal.Interface().(bool)
								}
							}
						}

						inputSelect.NewOption(label, value, disabled, selected)
					}
				}

			//

			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					// tagDataInternal.Func is the user function
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					// fieldVal.Interface() is the struct pointer that collects user data
					fieldVal.Interface(),
				}

				// explanation
				//   inputSelect.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
				//   The first element of the array is the user function
				//   From the second element onwards, they are internal functions and must be called after the user function in case the user has changed any value.
				inputSelect.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagData.Label),
		inputSelect,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Select{}) {
			r := fieldVal.Interface().(Select)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentRadio(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	inputDivRadio := factoryBrowser.NewTagDiv().
		Class("inputRadio").
		Data(map[string]string{"cell": "inputRadio"})

	elementOriginal := element
	radioComponent := Radio{}

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move the element from pointer to struct
	element = element.Elem()

	// Checks if the import of `components.Radio` was done
	if fieldRadio := element.FieldByName("Radio"); !fieldRadio.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Radio` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Radio"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         List *[]RadioData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label 3,value 3\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       type RadioData struct {"))
		err = errors.Join(err, fmt.Errorf("         TagRadio *html.TagInputRadio `wasmPanel:\"type:inputTagRadio\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         TagLabel *html.TagLabel      `wasmPanel:\"type:inputTagLabel\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Label    string              `wasmPanel:\"type:label\"`"))
		err = errors.Join(err, fmt.Errorf("         Value    string              `wasmPanel:\"type:value\"`"))
		err = errors.Join(err, fmt.Errorf("         Disabled bool                `wasmPanel:\"type:disabled\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Selected bool                `wasmPanel:\"type:selected\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Change   *RadioChange        `wasmPanel:\"type:listener;event:change;func:OnChangeEvent\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       // Note: Use `>` to set value as selected. ie. >label,value"))
		err = errors.Join(err, fmt.Errorf("       type RadioChange struct {"))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmGet:\"value\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *RadioChange) OnChangeEvent(event RadioChange, reference *Body) {"))
		err = errors.Join(err, fmt.Errorf("         log.Printf(\"value: %%v\", event.Value)"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Radio
		newInstance := reflect.New(fieldRadio.Type())
		fieldRadio.Set(newInstance.Elem())

		// Initializes the input tags within Radio
		//radioComponent.__radioTag = inputRadio // todo: fazer

		// __radioOnInputEvent is the pointer sent when the `change` event happens
		radioComponent.__change = new(__radioOnInputEvent)

		// populates the component.Radio within the user component
		componentRange := element.FieldByName("Radio")
		componentRange.Set(reflect.ValueOf(radioComponent))
	}

	err = e.verifyTypesComponentSelect(element) // todo: mudar este nome
	if err != nil {
		return
	}

	fieldNameInputTagLabel := ""
	fieldNameInputTagRadio := ""
	fieldNameLabel := ""
	fieldNameValue := ""
	fieldNameDisabled := ""
	fieldNameSelected := ""
	fieldNameListener := ""
	tagListener := new(tag)
	typeListener := reflect.StructField{}

	fieldTyp := element.Type()
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {
			case "value":

				// fieldVal.Interface() é *[]struct{...}, por isto .Elem().Elem(), ou *[] -> struct{...}
				fieldTyp := reflect.TypeOf(fieldVal.Interface()).Elem().Elem()
				for k := 0; k != fieldTyp.NumField(); k += 1 {
					fieldTyp := fieldTyp.Field(k)
					tagRaw := fieldTyp.Tag.Get("wasmPanel")
					if tagRaw != "" {
						tagDataInternal := new(tag)
						if err = tagDataInternal.init(tagRaw); err != nil {
							err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
							return
						}

						switch tagDataInternal.Type {
						case "inputTagLabel":
							fieldNameInputTagLabel = fieldTyp.Name
						case "inputTagRadio":
							fieldNameInputTagRadio = fieldTyp.Name
						case "label":
							fieldNameLabel = fieldTyp.Name
						case "value":
							fieldNameValue = fieldTyp.Name
						case "disabled":
							fieldNameDisabled = fieldTyp.Name
						case "selected":
							fieldNameSelected = fieldTyp.Name
						case "listener":
							fieldNameListener = fieldTyp.Name
							tagListener = tagDataInternal
							typeListener = fieldTyp
						}
					}
				}
			}
		}
	}

	//fieldTyp := element.Type()

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {
			//case "inputTagRadio":
			//	fieldVal.Set(reflect.ValueOf(inputRadio))

			case "value":

				fieldValPointer := fieldVal

				// pointer is not nil
				// Move the element from pointer to struct
				fieldVal = fieldVal.Elem()

				var inputLabel *html.TagLabel
				var inputRadio *html.TagInputRadio

				if fieldVal.IsZero() {

					var sliceValue reflect.Value
					var sliceType reflect.Type

					// fieldVal.Interface() é *[]struct{...}, por isto .Elem(), ou * -> []struct{...}
					sliceValue = reflect.ValueOf(fieldValPointer.Interface()).Elem()
					sliceType = reflect.TypeOf(sliceValue.Interface())
					newSlice := reflect.MakeSlice(sliceType, 0, 0)
					sliceValue.Set(newSlice)

					elemType := sliceType.Elem()
					newElem := reflect.New(elemType).Elem()

					if tagDataInternal.Default != "" {
						optionList := strings.Split(tagDataInternal.Default, ",")
						if len(optionList)%2 != 0 {
							err = fmt.Errorf("%v.%v: the correct format from tag value is: `wasmPanel:\"type:value;default:label1,value1,label2,value2,labelN,valueN\"`, where value and label, must be a pair", element.Type().Name(), fieldTyp.Name)
							return
						}

						for k := 0; k != len(optionList); k += 2 {

							inputRadio = factoryBrowser.NewTagInputRadio()
							inputLabel = factoryBrowser.NewTagLabel()

							// if label start with `>` the option is selected
							selected := false
							if strings.HasPrefix(optionList[k], ">") {
								optionList[k] = optionList[k][1:]
								selected = true
							}

							inputRadio.Value(optionList[k+1]).Disabled(false).Checked(selected).Class("inputRadio").Name(tagDataInternal.Name)
							inputLabel.Text(optionList[k]).Append(inputRadio)

							if fieldNameInputTagLabel != "" {
								newElem.FieldByName(fieldNameInputTagLabel).Set(reflect.ValueOf(inputLabel))
							}

							if fieldNameInputTagRadio != "" {
								newElem.FieldByName(fieldNameInputTagRadio).Set(reflect.ValueOf(inputRadio))
							}

							if fieldNameLabel != "" {
								newElem.FieldByName(fieldNameLabel).SetString(optionList[k+1])
							}

							if fieldNameValue != "" {
								newElem.FieldByName(fieldNameValue).SetString(optionList[k])
							}

							if fieldNameDisabled != "" {
								newElem.FieldByName(fieldNameDisabled).SetBool(false)
							}

							if fieldNameSelected != "" {
								newElem.FieldByName(fieldNameDisabled).SetBool(selected)
							}

							if fieldNameListener != "" {
								// The field must be a pointer, or it cannot be populated
								if typeListener.Type.Kind() != reflect.Pointer {
									log.Printf("error: %v.%v deve ser um ponteiro", newElem.Type().Name(), typeListener.Type.Name())
									return
								}

								if !typeListener.IsExported() {
									log.Printf("error: %v.%v não pode ser definido automaticamente.", newElem.Type().Name(), fieldNameListener)
									log.Printf("         isto geralmente acontece quando %v.%v não é público.", newElem.Type().Name(), fieldNameListener)
									return
								}

								newInstance := reflect.New(typeListener.Type.Elem())
								newElem.FieldByName(fieldNameListener).Set(newInstance)

								var methods []reflect.Value
								var params []interface{}

								// Passes the functions to be executed in the listener
								methods = []reflect.Value{
									// tagDataInternal.Func is the user function
									newElem.FieldByName(fieldNameListener).MethodByName(tagListener.Func),
								}

								// Pass variable pointers
								params = []interface{}{
									// fieldVal.Interface() is the struct pointer that collects user data
									newElem.FieldByName(fieldNameListener).Interface(),
								}

								inputRadio.ListenerAddReflect(tagListener.Event, params, methods, e.ref)
							}

							sliceValue.Set(reflect.Append(sliceValue, newElem))

							inputDivRadio.Append(
								factoryBrowser.NewTagSpan().AddStyle("flex", 1).Append(inputLabel),
							)
						}
					}

				} else {

					// run inside slice data
					for iField := 0; iField != fieldVal.Len(); iField += 1 {
						keyVal := fieldVal.Index(iField)

						inputRadio = factoryBrowser.NewTagInputRadio()
						inputLabel = factoryBrowser.NewTagLabel()

						// get label, value, disabled and selected
						var label, value string
						var disabled, selected bool
						for ik := 0; ik != keyVal.NumField(); ik += 1 {

							optionVal := keyVal.Field(ik)
							optionTyp := reflect.TypeOf(keyVal.Interface()).Field(ik)

							optionTagRaw := optionTyp.Tag.Get("wasmPanel")
							if optionTagRaw != "" {
								optionTag := new(tag)
								optionTag.init(optionTagRaw)

								switch optionTag.Type {
								case "inputTagLabel":
									optionVal.Set(reflect.ValueOf(inputLabel))
								case "inputTagRadio":
									optionVal.Set(reflect.ValueOf(inputRadio))
								case "label":
									label = optionVal.Interface().(string)
								case "value":
									value = optionVal.Interface().(string)
								case "disabled":
									disabled = optionVal.Interface().(bool)
								case "selected":
									selected = optionVal.Interface().(bool)
								case "listener":
									// The field must be a pointer, or it cannot be populated
									if typeListener.Type.Kind() != reflect.Pointer {
										log.Printf("error: %v.%v deve ser um ponteiro", optionVal.Type().Name(), typeListener.Type.Name())
										return
									}

									if !typeListener.IsExported() {
										log.Printf("error: %v.%v não pode ser definido automaticamente.", optionVal.Type().Name(), fieldNameListener)
										log.Printf("         isto geralmente acontece quando %v.%v não é público.", optionVal.Type().Name(), fieldNameListener)
										return
									}

									newInstance := reflect.New(typeListener.Type.Elem())
									keyVal.FieldByName(fieldNameListener).Set(newInstance)

									var methods []reflect.Value
									var params []interface{}

									// Passes the functions to be executed in the listener
									methods = []reflect.Value{
										// tagDataInternal.Func is the user function
										keyVal.FieldByName(fieldNameListener).MethodByName(tagListener.Func),
									}

									// Pass variable pointers
									params = []interface{}{
										// fieldVal.Interface() is the struct pointer that collects user data
										keyVal.FieldByName(fieldNameListener).Interface(),
									}

									inputRadio.ListenerAddReflect(tagListener.Event, params, methods, e.ref)
								}

							}
						}

						inputRadio.Value(value).Disabled(disabled).Checked(selected).Class("inputRadio").Name(tagDataInternal.Name)
						inputLabel.Text(label).Append(inputRadio)

						inputDivRadio.Append(
							factoryBrowser.NewTagSpan().AddStyle("flex", 1).Append(inputLabel),
						)

						//inputSelect.NewOption(label, value, disabled, selected)
					}
				}

			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagData.Label),
		inputDivRadio,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Radio{}) {
			r := fieldVal.Interface().(Radio)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentCheckbox(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	inputDivCheckbox := factoryBrowser.NewTagDiv().
		Class("inputCheckbox").
		Data(map[string]string{"cell": "inputCheckbox"})

	elementOriginal := element
	checkboxComponent := Checkbox{}

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move the element from pointer to struct
	element = element.Elem()

	// Checks if the import of `components.Checkbox` was done
	if fieldCheckbox := element.FieldByName("Checkbox"); !fieldCheckbox.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Checkbox` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Checkbox"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         List *[]CheckboxData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label 3,value 3\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       type CheckboxData struct {"))
		err = errors.Join(err, fmt.Errorf("         TagCheckbox *html.TagInputCheckbox `wasmPanel:\"type:inputTagCheckbox\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         TagLabel *html.TagLabel      `wasmPanel:\"type:inputTagLabel\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Label    string              `wasmPanel:\"type:label\"`"))
		err = errors.Join(err, fmt.Errorf("         Value    string              `wasmPanel:\"type:value\"`"))
		err = errors.Join(err, fmt.Errorf("         Disabled bool                `wasmPanel:\"type:disabled\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Selected bool                `wasmPanel:\"type:selected\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Change   *CheckboxChange        `wasmPanel:\"type:listener;event:change;func:OnChangeEvent\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       // Note: Use `>` to set value as selected. ie. >label,value"))
		err = errors.Join(err, fmt.Errorf("       type CheckboxChange struct {"))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmGet:\"value\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *CheckboxChange) OnChangeEvent(event CheckboxChange, reference *Body) {"))
		err = errors.Join(err, fmt.Errorf("         log.Printf(\"value: %%v\", event.Value)"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Checkbox
		newInstance := reflect.New(fieldCheckbox.Type())
		fieldCheckbox.Set(newInstance.Elem())

		// Initializes the input tags within Checkbox
		//checkboxComponent.__checkboxTag = inputCheckbox // todo: fazer

		// __checkboxOnInputEvent is the pointer sent when the `change` event happens
		checkboxComponent.__change = new(__checkboxOnInputEvent)

		// populates the component.Checkbox within the user component
		componentRange := element.FieldByName("Checkbox")
		componentRange.Set(reflect.ValueOf(checkboxComponent))
	}

	err = e.verifyTypesComponentSelect(element) // todo: mudar este nome
	if err != nil {
		return
	}

	fieldNameInputTagLabel := ""
	fieldNameInputTagCheckbox := ""
	fieldNameLabel := ""
	fieldNameValue := ""
	fieldNameDisabled := ""
	fieldNameSelected := ""
	fieldNameListener := ""
	tagListener := new(tag)
	typeListener := reflect.StructField{}

	fieldTyp := element.Type()
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {
			case "value":

				// fieldVal.Interface() é *[]struct{...}, por isto .Elem().Elem(), ou *[] -> struct{...}
				fieldTyp := reflect.TypeOf(fieldVal.Interface()).Elem().Elem()
				for k := 0; k != fieldTyp.NumField(); k += 1 {
					fieldTyp := fieldTyp.Field(k)
					tagRaw := fieldTyp.Tag.Get("wasmPanel")
					if tagRaw != "" {
						tagDataInternal := new(tag)
						if err = tagDataInternal.init(tagRaw); err != nil {
							err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
							return
						}

						switch tagDataInternal.Type {
						case "inputTagLabel":
							fieldNameInputTagLabel = fieldTyp.Name
						case "inputTagCheckbox":
							fieldNameInputTagCheckbox = fieldTyp.Name
						case "label":
							fieldNameLabel = fieldTyp.Name
						case "value":
							fieldNameValue = fieldTyp.Name
						case "disabled":
							fieldNameDisabled = fieldTyp.Name
						case "selected":
							fieldNameSelected = fieldTyp.Name
						case "listener":
							fieldNameListener = fieldTyp.Name
							tagListener = tagDataInternal
							typeListener = fieldTyp
						}
					}
				}
			}
		}
	}

	//fieldTyp := element.Type()

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			if err = tagDataInternal.init(tagRaw); err != nil {
				err = fmt.Errorf("error: the component %v has an error in one of the tags. the answer during processing was: %v", element.Type().Name(), err)
				return
			}

			switch tagDataInternal.Type {
			//case "inputTagCheckbox":
			//	fieldVal.Set(reflect.ValueOf(inputCheckbox))

			case "value":

				fieldValPointer := fieldVal

				// pointer is not nil
				// Move the element from pointer to struct
				fieldVal = fieldVal.Elem()

				var inputLabel *html.TagLabel
				var inputCheckbox *html.TagInputCheckBox

				if fieldVal.IsZero() {

					var sliceValue reflect.Value
					var sliceType reflect.Type

					// fieldVal.Interface() é *[]struct{...}, por isto .Elem(), ou * -> []struct{...}
					sliceValue = reflect.ValueOf(fieldValPointer.Interface()).Elem()
					sliceType = reflect.TypeOf(sliceValue.Interface())
					newSlice := reflect.MakeSlice(sliceType, 0, 0)
					sliceValue.Set(newSlice)

					elemType := sliceType.Elem()
					newElem := reflect.New(elemType).Elem()

					if tagDataInternal.Default != "" {
						optionList := strings.Split(tagDataInternal.Default, ",")
						if len(optionList)%2 != 0 {
							err = fmt.Errorf("%v.%v: the correct format from tag value is: `wasmPanel:\"type:value;default:label1,value1,label2,value2,labelN,valueN\"`, where value and label, must be a pair", element.Type().Name(), fieldTyp.Name)
							return
						}

						for k := 0; k != len(optionList); k += 2 {

							inputCheckbox = factoryBrowser.NewTagInputCheckBox()
							inputLabel = factoryBrowser.NewTagLabel()

							// if label start with `>` the option is selected
							selected := false
							if strings.HasPrefix(optionList[k], ">") {
								optionList[k] = optionList[k][1:]
								selected = true
							}

							inputCheckbox.Value(optionList[k+1]).Disabled(false).Checked(selected).Class("inputCheckbox").Name(tagDataInternal.Name)
							inputLabel.Text(optionList[k]).Append(inputCheckbox)

							if fieldNameInputTagLabel != "" {
								newElem.FieldByName(fieldNameInputTagLabel).Set(reflect.ValueOf(inputLabel))
							}

							if fieldNameInputTagCheckbox != "" {
								newElem.FieldByName(fieldNameInputTagCheckbox).Set(reflect.ValueOf(inputCheckbox))
							}

							if fieldNameLabel != "" {
								newElem.FieldByName(fieldNameLabel).SetString(optionList[k+1])
							}

							if fieldNameValue != "" {
								newElem.FieldByName(fieldNameValue).SetString(optionList[k])
							}

							if fieldNameDisabled != "" {
								newElem.FieldByName(fieldNameDisabled).SetBool(false)
							}

							if fieldNameSelected != "" {
								newElem.FieldByName(fieldNameDisabled).SetBool(selected)
							}

							if fieldNameListener != "" {
								// The field must be a pointer, or it cannot be populated
								if typeListener.Type.Kind() != reflect.Pointer {
									log.Printf("error: %v.%v deve ser um ponteiro", newElem.Type().Name(), typeListener.Type.Name())
									return
								}

								if !typeListener.IsExported() {
									log.Printf("error: %v.%v não pode ser definido automaticamente.", newElem.Type().Name(), fieldNameListener)
									log.Printf("         isto geralmente acontece quando %v.%v não é público.", newElem.Type().Name(), fieldNameListener)
									return
								}

								newInstance := reflect.New(typeListener.Type.Elem())
								newElem.FieldByName(fieldNameListener).Set(newInstance)

								var methods []reflect.Value
								var params []interface{}

								// Passes the functions to be executed in the listener
								methods = []reflect.Value{
									// tagDataInternal.Func is the user function
									newElem.FieldByName(fieldNameListener).MethodByName(tagListener.Func),
								}

								// Pass variable pointers
								params = []interface{}{
									// fieldVal.Interface() is the struct pointer that collects user data
									newElem.FieldByName(fieldNameListener).Interface(),
								}

								inputCheckbox.ListenerAddReflect(tagListener.Event, params, methods, e.ref)
							}

							sliceValue.Set(reflect.Append(sliceValue, newElem))

							inputDivCheckbox.Append(
								factoryBrowser.NewTagSpan().AddStyle("flex", 1).Append(inputLabel),
							)
						}
					}

				} else {

					// run inside slice data
					for iField := 0; iField != fieldVal.Len(); iField += 1 {
						keyVal := fieldVal.Index(iField)

						inputCheckbox = factoryBrowser.NewTagInputCheckBox()
						inputLabel = factoryBrowser.NewTagLabel()

						// get label, value, disabled and selected
						var label, value string
						var disabled, selected bool
						for ik := 0; ik != keyVal.NumField(); ik += 1 {

							optionVal := keyVal.Field(ik)
							optionTyp := reflect.TypeOf(keyVal.Interface()).Field(ik)

							optionTagRaw := optionTyp.Tag.Get("wasmPanel")
							if optionTagRaw != "" {
								optionTag := new(tag)
								optionTag.init(optionTagRaw)

								switch optionTag.Type {
								case "inputTagLabel":
									optionVal.Set(reflect.ValueOf(inputLabel))
								case "inputTagCheckbox":
									optionVal.Set(reflect.ValueOf(inputCheckbox))
								case "label":
									label = optionVal.Interface().(string)
								case "value":
									value = optionVal.Interface().(string)
								case "disabled":
									disabled = optionVal.Interface().(bool)
								case "selected":
									selected = optionVal.Interface().(bool)
								case "listener":
									// The field must be a pointer, or it cannot be populated
									if typeListener.Type.Kind() != reflect.Pointer {
										log.Printf("error: %v.%v deve ser um ponteiro", optionVal.Type().Name(), typeListener.Type.Name())
										return
									}

									if !typeListener.IsExported() {
										log.Printf("error: %v.%v não pode ser definido automaticamente.", optionVal.Type().Name(), fieldNameListener)
										log.Printf("         isto geralmente acontece quando %v.%v não é público.", optionVal.Type().Name(), fieldNameListener)
										return
									}

									newInstance := reflect.New(typeListener.Type.Elem())
									keyVal.FieldByName(fieldNameListener).Set(newInstance)

									var methods []reflect.Value
									var params []interface{}

									// Passes the functions to be executed in the listener
									methods = []reflect.Value{
										// tagDataInternal.Func is the user function
										keyVal.FieldByName(fieldNameListener).MethodByName(tagListener.Func),
									}

									// Pass variable pointers
									params = []interface{}{
										// fieldVal.Interface() is the struct pointer that collects user data
										keyVal.FieldByName(fieldNameListener).Interface(),
									}

									inputCheckbox.ListenerAddReflect(tagListener.Event, params, methods, e.ref)
								}

							}
						}

						inputCheckbox.Value(value).Disabled(disabled).Checked(selected).Class("inputCheckbox").Name(tagDataInternal.Name)
						inputLabel.Text(label).Append(inputCheckbox)

						inputDivCheckbox.Append(
							factoryBrowser.NewTagSpan().AddStyle("flex", 1).Append(inputLabel),
						)

						//inputSelect.NewOption(label, value, disabled, selected)
					}
				}

			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().AddStyle("flex", 1).Text(tagData.Label),
		inputDivCheckbox,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Checkbox{}) {
			r := fieldVal.Interface().(Checkbox)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processHeaderTextAddDragListener(dragIcon *html.TagDiv) {
	dragIcon.Get().Call("addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) any {
		e.isDragging = true
		e.offsetX = args[0].Get("clientX").Int() - e.panelFather.Get().Call("getBoundingClientRect").Get("left").Int()
		e.offsetY = args[0].Get("clientY").Int() - e.panelFather.Get().Call("getBoundingClientRect").Get("top").Int()
		return nil
	}))

	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) any {
		if e.isDragging {
			e.panelFather.AddStyle("left", fmt.Sprintf("%vpx", args[0].Get("clientX").Int()-e.offsetX))
			e.panelFather.AddStyle("top", fmt.Sprintf("%vpx", args[0].Get("clientY").Int()-e.offsetY))
		}
		return nil
	}))

	js.Global().Get("document").Call("addEventListener", "mouseup", js.FuncOf(func(this js.Value, args []js.Value) any {
		e.isDragging = false
		return nil
	}))
}

func (e *Components) processHeaderTextAddMinimizeListener(minimizeIcon *html.TagDiv) {
	minimizeIcon.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		args[0].Call("stopPropagation")

		display := "none"
		if e.minimized {
			display = "block"
		}
		e.minimized = !e.minimized
		e.panelBody.AddStyle("display", display)

		return nil
	}))
}

func (e *Components) processHeaderTextAddCloseListener(closeIcon *html.TagDiv) {
	closeIcon.Get().Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		args[0].Call("stopPropagation")

		e.panelFather.Fade(200 * time.Millisecond)
		go func() {
			time.Sleep(2 * time.Second)
			e.panelFather.Fade(200 * time.Millisecond)
		}()
		return nil
	}))
}

func (e *Components) Show() {
	e.panelFather.Get().Get("classList").Call("toggle", "open") // todo: se tipo fechar
	e.panelFather.AddStyle("display", "visible")
}

func (e *Components) processHeaderText(element reflect.Value, defaultText string, father *html.TagDiv) (dragIcon, minimizeIcon, closeIcon *html.TagDiv) {

	var ok bool
	var text string

	text, ok = element.Interface().(string)
	if ok && text != "" {
		defaultText = text
	}

	// <div class="panelHeader">
	//   <div class="headerText">Panel</div>
	//   <div class="dragIcon"></div>
	//   <div class="closeIconPanel">ˇ</div>
	// </div>

	// minimize(ˇ) close(⤬)
	closeIcon = factoryBrowser.NewTagDiv().AddStyle("cursor", "pointer").Html("⊗")
	dragIcon = factoryBrowser.NewTagDiv().AddStyle("cursor", "move").Html("◇")
	minimizeIcon = factoryBrowser.NewTagDiv().AddStyle("cursor", "pointer").Html("▾")
	panelHeader := factoryBrowser.NewTagDiv().Class("panelHeader").Append(
		factoryBrowser.NewTagDiv().
			Class("headerText").
			AddStyle("fontWeight", "bold").
			AddStyle("flex", 1).
			//AddStyle("textAlign", "center").
			Html(defaultText),
		dragIcon,
		factoryBrowser.NewTagDiv().Html("&nbsp;"),
		minimizeIcon,
		factoryBrowser.NewTagDiv().Html("&nbsp;"),
		closeIcon,
	)
	father.Append(
		panelHeader,
	)

	for k := range e.setupPanelHeader {
		panelHeader.AddStyle(k, e.setupPanelHeader[k])
	}

	return
}
