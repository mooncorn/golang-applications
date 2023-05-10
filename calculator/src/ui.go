package src

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UI struct {
	TerminalApp *tview.Application
	CmdInputs   chan uicommand
	HistoryBox  *tview.TextView
	ResultBox   *tview.TextView
	*Calculator
}

type uicommand struct {
	cmdtype string
	cmdarg1 string
	cmdarg2 string
}

func NewUI(calculator *Calculator) *UI {
	app := tview.NewApplication()

	cmdchan := make(chan uicommand)

	titlebox := tview.NewTextView().
		SetText("Calculator").
		SetTextColor(tcell.ColorWhite).
		SetTextAlign(tview.AlignCenter)

	titlebox.SetBorder(true).SetBorderColor(tcell.ColorGreen)

	historyBox := tview.NewTextView().SetDynamicColors(true).SetChangedFunc(func() {
		app.Draw()
	})

	historyBox.SetBorder(true).SetBorderColor(tcell.ColorGreen).
		SetTitle("Logs").
		SetTitleAlign(tview.AlignLeft).SetTitleColor(tcell.ColorWhite)

	usage := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[red]/add[orange] <n1> [n2] [yellow]| [red]/substract[orange] <n1> [n2] [yellow]| [red]/multiply[orange] <n1> [n2] [yellow]| [red]/divide[orange] <n1> [n2] [yellow]| [red]/clear [yellow]| [red]/quit`)
	usage.SetBorder(true).SetBorderColor(tcell.ColorGreen).
		SetTitle("Usage").SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorWhite).SetBorderPadding(0, 0, 1, 0)

	resultBox := tview.NewTextView()
	resultBox.SetBorder(true).SetBorderColor(tcell.ColorGreen).
		SetTitle("Result").SetTitleAlign(tview.AlignLeft).SetTitleColor(tcell.ColorWhite)

	fmt.Fprintf(resultBox, "%f", 0.0)

	input := tview.NewInputField().SetLabel(">").
		SetLabelColor(tcell.ColorGreen).SetFieldWidth(0).SetFieldBackgroundColor(tcell.ColorBlack)

	input.SetBorder(true).SetBorderColor(tcell.ColorGreen).
		SetTitle("Input").SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorWhite).SetBorderPadding(0, 0, 1, 0)

	input.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}
		line := input.GetText()
		if len(line) == 0 {
			return
		}
		if strings.HasPrefix(line, "/") {
			cmdparts := strings.Split(line, " ")

			if len(cmdparts) == 1 {
				cmdparts = append(cmdparts, "")
				cmdparts = append(cmdparts, "")
			} else if len(cmdparts) == 2 {
				cmdparts = append(cmdparts, "")
			}
			cmdchan <- uicommand{cmdtype: cmdparts[0], cmdarg1: cmdparts[1], cmdarg2: cmdparts[2]}
		} else {
			calculator.Logs <- callog{logprefix: "badcmd",
				logmsg: fmt.Sprintf("Unsupported command %s", line)}
		}
		input.SetText("")
	})

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(titlebox, 3, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(historyBox, 0, 1, false).
			AddItem(resultBox, 20, 1, false),
			0, 8, false).AddItem(input, 3, 1, true).
		AddItem(usage, 3, 1, false)

	app.SetRoot(flex, true)

	return &UI{
		TerminalApp: app,
		HistoryBox:  historyBox,
		ResultBox:   resultBox,
		CmdInputs:   cmdchan,
		Calculator:  calculator,
	}
}

func (ui *UI) Run() error {
	go ui.starteventhandler()
	return ui.TerminalApp.Run()
}

func (ui *UI) starteventhandler() {
	refreshticker := time.NewTicker(time.Second)
	defer refreshticker.Stop()
	for {
		select {
		case cmd := <-ui.CmdInputs:
			go ui.handleCommand(cmd)
		case log := <-ui.Logs:
			ui.display_logMessage(log)
		case op := <-ui.Operations:
			ui.display_operationMessage(op)
			ui.display_resultMessage(ui.Result)
		}
	}
}

func (ui *UI) handleCommand(cmd uicommand) {
	switch cmd.cmdtype {
	case "/quit":
		ui.TerminalApp.Stop()
		return
	case "/clear":
		ui.HistoryBox.Clear()
		ui.display_resultMessage(0)
		ui.Result = 0
	case "/add", "/substract", "/divide", "/multiply":
		if len(cmd.cmdarg1) > 0 && len(cmd.cmdarg2) > 0 {
			num1, err1 := strconv.ParseFloat(cmd.cmdarg1, 64)
			num2, err2 := strconv.ParseFloat(cmd.cmdarg2, 64)
			if err1 != nil && err2 != nil {
				ui.Logs <- callog{logprefix: "badarg", logmsg: "args must be numbers"}
			} else {
				ui.ExecuteOperation(cmd.cmdtype[1:], num1, num2)
			}
		} else if len(cmd.cmdarg1) > 0 {
			num, err := strconv.ParseFloat(cmd.cmdarg1, 64)
			if err != nil {
				ui.Logs <- callog{logprefix: "badarg", logmsg: "arg1 has to be a number"}
			} else {
				ui.ExecuteOperation(cmd.cmdtype[1:], ui.Result, num)
			}
		} else {
			ui.Logs <- callog{logprefix: "badcmd", logmsg: "missing arg1 for the operation"}
		}
	default:
		ui.Logs <- callog{logprefix: "badcmd",
			logmsg: fmt.Sprintf("Unsupported command %s", cmd.cmdtype)}
	}
}

func (ui *UI) display_resultMessage(res float64) {
	ui.ResultBox.Clear()
	fmt.Fprintf(ui.ResultBox, "%s", strconv.FormatFloat(res, 'f', -1, 64))
}

func (ui *UI) display_operationMessage(op operation) {
	prompt := fmt.Sprintf("[green]<%s>:[-]", "Operation")
	fmt.Fprintf(ui.HistoryBox, "%s %s %s %s\n", prompt, strconv.FormatFloat(op.num1, 'f', -1, 64), op.operation, strconv.FormatFloat(op.num2, 'f', -1, 64))
}

func (ui *UI) display_logMessage(log callog) {
	prompt := fmt.Sprintf("[yellow]<%s>:[-]", "Log")
	fmt.Fprintf(ui.HistoryBox, "%s %s\n", prompt, log.logmsg)
}
