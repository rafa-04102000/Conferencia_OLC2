// Code generated from Gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr4-go/antlr/v4"

// BaseGramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type BaseGramaticaListener struct{}

var _ GramaticaListener = &BaseGramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGramaticaListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGramaticaListener) ExitProgram(ctx *ProgramContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGramaticaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGramaticaListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGramaticaListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGramaticaListener) ExitStatement(ctx *StatementContext) {}

// EnterMainfunction is called when production mainfunction is entered.
func (s *BaseGramaticaListener) EnterMainfunction(ctx *MainfunctionContext) {}

// ExitMainfunction is called when production mainfunction is exited.
func (s *BaseGramaticaListener) ExitMainfunction(ctx *MainfunctionContext) {}

// EnterDeclarationExplicitVar is called when production DeclarationExplicitVar is entered.
func (s *BaseGramaticaListener) EnterDeclarationExplicitVar(ctx *DeclarationExplicitVarContext) {}

// ExitDeclarationExplicitVar is called when production DeclarationExplicitVar is exited.
func (s *BaseGramaticaListener) ExitDeclarationExplicitVar(ctx *DeclarationExplicitVarContext) {}

// EnterDeclarationImplicitVar is called when production DeclarationImplicitVar is entered.
func (s *BaseGramaticaListener) EnterDeclarationImplicitVar(ctx *DeclarationImplicitVarContext) {}

// ExitDeclarationImplicitVar is called when production DeclarationImplicitVar is exited.
func (s *BaseGramaticaListener) ExitDeclarationImplicitVar(ctx *DeclarationImplicitVarContext) {}

// EnterDeclarationExplicitSlice is called when production DeclarationExplicitSlice is entered.
func (s *BaseGramaticaListener) EnterDeclarationExplicitSlice(ctx *DeclarationExplicitSliceContext) {}

// ExitDeclarationExplicitSlice is called when production DeclarationExplicitSlice is exited.
func (s *BaseGramaticaListener) ExitDeclarationExplicitSlice(ctx *DeclarationExplicitSliceContext) {}

// EnterDeclarationImplicitSlice is called when production DeclarationImplicitSlice is entered.
func (s *BaseGramaticaListener) EnterDeclarationImplicitSlice(ctx *DeclarationImplicitSliceContext) {}

// ExitDeclarationImplicitSlice is called when production DeclarationImplicitSlice is exited.
func (s *BaseGramaticaListener) ExitDeclarationImplicitSlice(ctx *DeclarationImplicitSliceContext) {}

// EnterStructDeclaration is called when production StructDeclaration is entered.
func (s *BaseGramaticaListener) EnterStructDeclaration(ctx *StructDeclarationContext) {}

// ExitStructDeclaration is called when production StructDeclaration is exited.
func (s *BaseGramaticaListener) ExitStructDeclaration(ctx *StructDeclarationContext) {}

// EnterStructInstanceDeclaration is called when production StructInstanceDeclaration is entered.
func (s *BaseGramaticaListener) EnterStructInstanceDeclaration(ctx *StructInstanceDeclarationContext) {
}

// ExitStructInstanceDeclaration is called when production StructInstanceDeclaration is exited.
func (s *BaseGramaticaListener) ExitStructInstanceDeclaration(ctx *StructInstanceDeclarationContext) {
}

// EnterFunctionDeclaration is called when production FunctionDeclaration is entered.
func (s *BaseGramaticaListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production FunctionDeclaration is exited.
func (s *BaseGramaticaListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseGramaticaListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseGramaticaListener) ExitParameters(ctx *ParametersContext) {}

// EnterNormalTypeParameter is called when production NormalTypeParameter is entered.
func (s *BaseGramaticaListener) EnterNormalTypeParameter(ctx *NormalTypeParameterContext) {}

// ExitNormalTypeParameter is called when production NormalTypeParameter is exited.
func (s *BaseGramaticaListener) ExitNormalTypeParameter(ctx *NormalTypeParameterContext) {}

// EnterSliceTypeParameter is called when production SliceTypeParameter is entered.
func (s *BaseGramaticaListener) EnterSliceTypeParameter(ctx *SliceTypeParameterContext) {}

// ExitSliceTypeParameter is called when production SliceTypeParameter is exited.
func (s *BaseGramaticaListener) ExitSliceTypeParameter(ctx *SliceTypeParameterContext) {}

// EnterTiposSlice is called when production tiposSlice is entered.
func (s *BaseGramaticaListener) EnterTiposSlice(ctx *TiposSliceContext) {}

// ExitTiposSlice is called when production tiposSlice is exited.
func (s *BaseGramaticaListener) ExitTiposSlice(ctx *TiposSliceContext) {}

// EnterStruct_field is called when production struct_field is entered.
func (s *BaseGramaticaListener) EnterStruct_field(ctx *Struct_fieldContext) {}

// ExitStruct_field is called when production struct_field is exited.
func (s *BaseGramaticaListener) ExitStruct_field(ctx *Struct_fieldContext) {}

// EnterLstStruct is called when production lstStruct is entered.
func (s *BaseGramaticaListener) EnterLstStruct(ctx *LstStructContext) {}

// ExitLstStruct is called when production lstStruct is exited.
func (s *BaseGramaticaListener) ExitLstStruct(ctx *LstStructContext) {}

// EnterListSlice is called when production listSlice is entered.
func (s *BaseGramaticaListener) EnterListSlice(ctx *ListSliceContext) {}

// ExitListSlice is called when production listSlice is exited.
func (s *BaseGramaticaListener) ExitListSlice(ctx *ListSliceContext) {}

// EnterAssignmentExpr is called when production AssignmentExpr is entered.
func (s *BaseGramaticaListener) EnterAssignmentExpr(ctx *AssignmentExprContext) {}

// ExitAssignmentExpr is called when production AssignmentExpr is exited.
func (s *BaseGramaticaListener) ExitAssignmentExpr(ctx *AssignmentExprContext) {}

// EnterAssignmentPlusExpr is called when production AssignmentPlusExpr is entered.
func (s *BaseGramaticaListener) EnterAssignmentPlusExpr(ctx *AssignmentPlusExprContext) {}

// ExitAssignmentPlusExpr is called when production AssignmentPlusExpr is exited.
func (s *BaseGramaticaListener) ExitAssignmentPlusExpr(ctx *AssignmentPlusExprContext) {}

// EnterAssignmentMinusExpr is called when production AssignmentMinusExpr is entered.
func (s *BaseGramaticaListener) EnterAssignmentMinusExpr(ctx *AssignmentMinusExprContext) {}

// ExitAssignmentMinusExpr is called when production AssignmentMinusExpr is exited.
func (s *BaseGramaticaListener) ExitAssignmentMinusExpr(ctx *AssignmentMinusExprContext) {}

// EnterAssignmentIncrementExpr is called when production AssignmentIncrementExpr is entered.
func (s *BaseGramaticaListener) EnterAssignmentIncrementExpr(ctx *AssignmentIncrementExprContext) {}

// ExitAssignmentIncrementExpr is called when production AssignmentIncrementExpr is exited.
func (s *BaseGramaticaListener) ExitAssignmentIncrementExpr(ctx *AssignmentIncrementExprContext) {}

// EnterAssignmentDecrementExpr is called when production AssignmentDecrementExpr is entered.
func (s *BaseGramaticaListener) EnterAssignmentDecrementExpr(ctx *AssignmentDecrementExprContext) {}

// ExitAssignmentDecrementExpr is called when production AssignmentDecrementExpr is exited.
func (s *BaseGramaticaListener) ExitAssignmentDecrementExpr(ctx *AssignmentDecrementExprContext) {}

// EnterSliceAssignmentExpr is called when production SliceAssignmentExpr is entered.
func (s *BaseGramaticaListener) EnterSliceAssignmentExpr(ctx *SliceAssignmentExprContext) {}

// ExitSliceAssignmentExpr is called when production SliceAssignmentExpr is exited.
func (s *BaseGramaticaListener) ExitSliceAssignmentExpr(ctx *SliceAssignmentExprContext) {}

// EnterAttributeAssignmentExpr is called when production AttributeAssignmentExpr is entered.
func (s *BaseGramaticaListener) EnterAttributeAssignmentExpr(ctx *AttributeAssignmentExprContext) {}

// ExitAttributeAssignmentExpr is called when production AttributeAssignmentExpr is exited.
func (s *BaseGramaticaListener) ExitAttributeAssignmentExpr(ctx *AttributeAssignmentExprContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseGramaticaListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseGramaticaListener) ExitPrint(ctx *PrintContext) {}

// EnterPrintList is called when production printList is entered.
func (s *BaseGramaticaListener) EnterPrintList(ctx *PrintListContext) {}

// ExitPrintList is called when production printList is exited.
func (s *BaseGramaticaListener) ExitPrintList(ctx *PrintListContext) {}

// EnterIf is called when production if is entered.
func (s *BaseGramaticaListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseGramaticaListener) ExitIf(ctx *IfContext) {}

// EnterElseIfBlock is called when production elseIfBlock is entered.
func (s *BaseGramaticaListener) EnterElseIfBlock(ctx *ElseIfBlockContext) {}

// ExitElseIfBlock is called when production elseIfBlock is exited.
func (s *BaseGramaticaListener) ExitElseIfBlock(ctx *ElseIfBlockContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseGramaticaListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseGramaticaListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterSwitch is called when production switch is entered.
func (s *BaseGramaticaListener) EnterSwitch(ctx *SwitchContext) {}

// ExitSwitch is called when production switch is exited.
func (s *BaseGramaticaListener) ExitSwitch(ctx *SwitchContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseGramaticaListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseGramaticaListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterDefaultBlock is called when production defaultBlock is entered.
func (s *BaseGramaticaListener) EnterDefaultBlock(ctx *DefaultBlockContext) {}

// ExitDefaultBlock is called when production defaultBlock is exited.
func (s *BaseGramaticaListener) ExitDefaultBlock(ctx *DefaultBlockContext) {}

// EnterForSimple is called when production ForSimple is entered.
func (s *BaseGramaticaListener) EnterForSimple(ctx *ForSimpleContext) {}

// ExitForSimple is called when production ForSimple is exited.
func (s *BaseGramaticaListener) ExitForSimple(ctx *ForSimpleContext) {}

// EnterForDeclaration is called when production ForDeclaration is entered.
func (s *BaseGramaticaListener) EnterForDeclaration(ctx *ForDeclarationContext) {}

// ExitForDeclaration is called when production ForDeclaration is exited.
func (s *BaseGramaticaListener) ExitForDeclaration(ctx *ForDeclarationContext) {}

// EnterForInSlice is called when production ForInSlice is entered.
func (s *BaseGramaticaListener) EnterForInSlice(ctx *ForInSliceContext) {}

// ExitForInSlice is called when production ForInSlice is exited.
func (s *BaseGramaticaListener) ExitForInSlice(ctx *ForInSliceContext) {}

// EnterIndividualBlock is called when production individualBlock is entered.
func (s *BaseGramaticaListener) EnterIndividualBlock(ctx *IndividualBlockContext) {}

// ExitIndividualBlock is called when production individualBlock is exited.
func (s *BaseGramaticaListener) ExitIndividualBlock(ctx *IndividualBlockContext) {}

// EnterBreakStatement is called when production BreakStatement is entered.
func (s *BaseGramaticaListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production BreakStatement is exited.
func (s *BaseGramaticaListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production ContinueStatement is entered.
func (s *BaseGramaticaListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production ContinueStatement is exited.
func (s *BaseGramaticaListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production ReturnStatement is entered.
func (s *BaseGramaticaListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production ReturnStatement is exited.
func (s *BaseGramaticaListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseGramaticaListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseGramaticaListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGramaticaListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGramaticaListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseGramaticaListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseGramaticaListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterIndexOfFunction is called when production IndexOfFunction is entered.
func (s *BaseGramaticaListener) EnterIndexOfFunction(ctx *IndexOfFunctionContext) {}

// ExitIndexOfFunction is called when production IndexOfFunction is exited.
func (s *BaseGramaticaListener) ExitIndexOfFunction(ctx *IndexOfFunctionContext) {}

// EnterJoinFunction is called when production JoinFunction is entered.
func (s *BaseGramaticaListener) EnterJoinFunction(ctx *JoinFunctionContext) {}

// ExitJoinFunction is called when production JoinFunction is exited.
func (s *BaseGramaticaListener) ExitJoinFunction(ctx *JoinFunctionContext) {}

// EnterLengthFunction is called when production LengthFunction is entered.
func (s *BaseGramaticaListener) EnterLengthFunction(ctx *LengthFunctionContext) {}

// ExitLengthFunction is called when production LengthFunction is exited.
func (s *BaseGramaticaListener) ExitLengthFunction(ctx *LengthFunctionContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseGramaticaListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseGramaticaListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterFloatExpr is called when production FloatExpr is entered.
func (s *BaseGramaticaListener) EnterFloatExpr(ctx *FloatExprContext) {}

// ExitFloatExpr is called when production FloatExpr is exited.
func (s *BaseGramaticaListener) ExitFloatExpr(ctx *FloatExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseGramaticaListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseGramaticaListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterDerefExpr is called when production DerefExpr is entered.
func (s *BaseGramaticaListener) EnterDerefExpr(ctx *DerefExprContext) {}

// ExitDerefExpr is called when production DerefExpr is exited.
func (s *BaseGramaticaListener) ExitDerefExpr(ctx *DerefExprContext) {}

// EnterParseFloatExpr is called when production ParseFloatExpr is entered.
func (s *BaseGramaticaListener) EnterParseFloatExpr(ctx *ParseFloatExprContext) {}

// ExitParseFloatExpr is called when production ParseFloatExpr is exited.
func (s *BaseGramaticaListener) ExitParseFloatExpr(ctx *ParseFloatExprContext) {}

// EnterNegExpr is called when production NegExpr is entered.
func (s *BaseGramaticaListener) EnterNegExpr(ctx *NegExprContext) {}

// ExitNegExpr is called when production NegExpr is exited.
func (s *BaseGramaticaListener) ExitNegExpr(ctx *NegExprContext) {}

// EnterAppendExpr is called when production AppendExpr is entered.
func (s *BaseGramaticaListener) EnterAppendExpr(ctx *AppendExprContext) {}

// ExitAppendExpr is called when production AppendExpr is exited.
func (s *BaseGramaticaListener) ExitAppendExpr(ctx *AppendExprContext) {}

// EnterOpExpr is called when production OpExpr is entered.
func (s *BaseGramaticaListener) EnterOpExpr(ctx *OpExprContext) {}

// ExitOpExpr is called when production OpExpr is exited.
func (s *BaseGramaticaListener) ExitOpExpr(ctx *OpExprContext) {}

// EnterSlicePrimitiveValue is called when production SlicePrimitiveValue is entered.
func (s *BaseGramaticaListener) EnterSlicePrimitiveValue(ctx *SlicePrimitiveValueContext) {}

// ExitSlicePrimitiveValue is called when production SlicePrimitiveValue is exited.
func (s *BaseGramaticaListener) ExitSlicePrimitiveValue(ctx *SlicePrimitiveValueContext) {}

// EnterTypeOfExpr is called when production TypeOfExpr is entered.
func (s *BaseGramaticaListener) EnterTypeOfExpr(ctx *TypeOfExprContext) {}

// ExitTypeOfExpr is called when production TypeOfExpr is exited.
func (s *BaseGramaticaListener) ExitTypeOfExpr(ctx *TypeOfExprContext) {}

// EnterStructAccess is called when production StructAccess is entered.
func (s *BaseGramaticaListener) EnterStructAccess(ctx *StructAccessContext) {}

// ExitStructAccess is called when production StructAccess is exited.
func (s *BaseGramaticaListener) ExitStructAccess(ctx *StructAccessContext) {}

// EnterFunctionCallExpr is called when production FunctionCallExpr is entered.
func (s *BaseGramaticaListener) EnterFunctionCallExpr(ctx *FunctionCallExprContext) {}

// ExitFunctionCallExpr is called when production FunctionCallExpr is exited.
func (s *BaseGramaticaListener) ExitFunctionCallExpr(ctx *FunctionCallExprContext) {}

// EnterAtoiExpr is called when production AtoiExpr is entered.
func (s *BaseGramaticaListener) EnterAtoiExpr(ctx *AtoiExprContext) {}

// ExitAtoiExpr is called when production AtoiExpr is exited.
func (s *BaseGramaticaListener) ExitAtoiExpr(ctx *AtoiExprContext) {}

// EnterParExpr is called when production ParExpr is entered.
func (s *BaseGramaticaListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production ParExpr is exited.
func (s *BaseGramaticaListener) ExitParExpr(ctx *ParExprContext) {}

// EnterStrExpr is called when production StrExpr is entered.
func (s *BaseGramaticaListener) EnterStrExpr(ctx *StrExprContext) {}

// ExitStrExpr is called when production StrExpr is exited.
func (s *BaseGramaticaListener) ExitStrExpr(ctx *StrExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseGramaticaListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseGramaticaListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseGramaticaListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseGramaticaListener) ExitIntExpr(ctx *IntExprContext) {}

// EnterRefExpr is called when production RefExpr is entered.
func (s *BaseGramaticaListener) EnterRefExpr(ctx *RefExprContext) {}

// ExitRefExpr is called when production RefExpr is exited.
func (s *BaseGramaticaListener) ExitRefExpr(ctx *RefExprContext) {}

// EnterSliceAccess is called when production SliceAccess is entered.
func (s *BaseGramaticaListener) EnterSliceAccess(ctx *SliceAccessContext) {}

// ExitSliceAccess is called when production SliceAccess is exited.
func (s *BaseGramaticaListener) ExitSliceAccess(ctx *SliceAccessContext) {}
