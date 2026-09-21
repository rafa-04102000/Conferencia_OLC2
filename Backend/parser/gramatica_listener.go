// Code generated from Gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr4-go/antlr/v4"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterMainfunction is called when entering the mainfunction production.
	EnterMainfunction(c *MainfunctionContext)

	// EnterDeclarationExplicitVar is called when entering the DeclarationExplicitVar production.
	EnterDeclarationExplicitVar(c *DeclarationExplicitVarContext)

	// EnterDeclarationImplicitVar is called when entering the DeclarationImplicitVar production.
	EnterDeclarationImplicitVar(c *DeclarationImplicitVarContext)

	// EnterDeclarationExplicitSlice is called when entering the DeclarationExplicitSlice production.
	EnterDeclarationExplicitSlice(c *DeclarationExplicitSliceContext)

	// EnterDeclarationImplicitSlice is called when entering the DeclarationImplicitSlice production.
	EnterDeclarationImplicitSlice(c *DeclarationImplicitSliceContext)

	// EnterStructDeclaration is called when entering the StructDeclaration production.
	EnterStructDeclaration(c *StructDeclarationContext)

	// EnterStructInstanceDeclaration is called when entering the StructInstanceDeclaration production.
	EnterStructInstanceDeclaration(c *StructInstanceDeclarationContext)

	// EnterFunctionDeclaration is called when entering the FunctionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterNormalTypeParameter is called when entering the NormalTypeParameter production.
	EnterNormalTypeParameter(c *NormalTypeParameterContext)

	// EnterSliceTypeParameter is called when entering the SliceTypeParameter production.
	EnterSliceTypeParameter(c *SliceTypeParameterContext)

	// EnterTiposSlice is called when entering the tiposSlice production.
	EnterTiposSlice(c *TiposSliceContext)

	// EnterStruct_field is called when entering the struct_field production.
	EnterStruct_field(c *Struct_fieldContext)

	// EnterLstStruct is called when entering the lstStruct production.
	EnterLstStruct(c *LstStructContext)

	// EnterListSlice is called when entering the listSlice production.
	EnterListSlice(c *ListSliceContext)

	// EnterAssignmentExpr is called when entering the AssignmentExpr production.
	EnterAssignmentExpr(c *AssignmentExprContext)

	// EnterAssignmentPlusExpr is called when entering the AssignmentPlusExpr production.
	EnterAssignmentPlusExpr(c *AssignmentPlusExprContext)

	// EnterAssignmentMinusExpr is called when entering the AssignmentMinusExpr production.
	EnterAssignmentMinusExpr(c *AssignmentMinusExprContext)

	// EnterAssignmentIncrementExpr is called when entering the AssignmentIncrementExpr production.
	EnterAssignmentIncrementExpr(c *AssignmentIncrementExprContext)

	// EnterAssignmentDecrementExpr is called when entering the AssignmentDecrementExpr production.
	EnterAssignmentDecrementExpr(c *AssignmentDecrementExprContext)

	// EnterSliceAssignmentExpr is called when entering the SliceAssignmentExpr production.
	EnterSliceAssignmentExpr(c *SliceAssignmentExprContext)

	// EnterAttributeAssignmentExpr is called when entering the AttributeAssignmentExpr production.
	EnterAttributeAssignmentExpr(c *AttributeAssignmentExprContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterPrintList is called when entering the printList production.
	EnterPrintList(c *PrintListContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterElseIfBlock is called when entering the elseIfBlock production.
	EnterElseIfBlock(c *ElseIfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterSwitch is called when entering the switch production.
	EnterSwitch(c *SwitchContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterDefaultBlock is called when entering the defaultBlock production.
	EnterDefaultBlock(c *DefaultBlockContext)

	// EnterForSimple is called when entering the ForSimple production.
	EnterForSimple(c *ForSimpleContext)

	// EnterForDeclaration is called when entering the ForDeclaration production.
	EnterForDeclaration(c *ForDeclarationContext)

	// EnterForInSlice is called when entering the ForInSlice production.
	EnterForInSlice(c *ForInSliceContext)

	// EnterIndividualBlock is called when entering the individualBlock production.
	EnterIndividualBlock(c *IndividualBlockContext)

	// EnterBreakStatement is called when entering the BreakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the ContinueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the ReturnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterIndexOfFunction is called when entering the IndexOfFunction production.
	EnterIndexOfFunction(c *IndexOfFunctionContext)

	// EnterJoinFunction is called when entering the JoinFunction production.
	EnterJoinFunction(c *JoinFunctionContext)

	// EnterLengthFunction is called when entering the LengthFunction production.
	EnterLengthFunction(c *LengthFunctionContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterFloatExpr is called when entering the FloatExpr production.
	EnterFloatExpr(c *FloatExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterDerefExpr is called when entering the DerefExpr production.
	EnterDerefExpr(c *DerefExprContext)

	// EnterParseFloatExpr is called when entering the ParseFloatExpr production.
	EnterParseFloatExpr(c *ParseFloatExprContext)

	// EnterNegExpr is called when entering the NegExpr production.
	EnterNegExpr(c *NegExprContext)

	// EnterAppendExpr is called when entering the AppendExpr production.
	EnterAppendExpr(c *AppendExprContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// EnterSlicePrimitiveValue is called when entering the SlicePrimitiveValue production.
	EnterSlicePrimitiveValue(c *SlicePrimitiveValueContext)

	// EnterTypeOfExpr is called when entering the TypeOfExpr production.
	EnterTypeOfExpr(c *TypeOfExprContext)

	// EnterStructAccess is called when entering the StructAccess production.
	EnterStructAccess(c *StructAccessContext)

	// EnterFunctionCallExpr is called when entering the FunctionCallExpr production.
	EnterFunctionCallExpr(c *FunctionCallExprContext)

	// EnterAtoiExpr is called when entering the AtoiExpr production.
	EnterAtoiExpr(c *AtoiExprContext)

	// EnterParExpr is called when entering the ParExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterStrExpr is called when entering the StrExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterRefExpr is called when entering the RefExpr production.
	EnterRefExpr(c *RefExprContext)

	// EnterSliceAccess is called when entering the SliceAccess production.
	EnterSliceAccess(c *SliceAccessContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitMainfunction is called when exiting the mainfunction production.
	ExitMainfunction(c *MainfunctionContext)

	// ExitDeclarationExplicitVar is called when exiting the DeclarationExplicitVar production.
	ExitDeclarationExplicitVar(c *DeclarationExplicitVarContext)

	// ExitDeclarationImplicitVar is called when exiting the DeclarationImplicitVar production.
	ExitDeclarationImplicitVar(c *DeclarationImplicitVarContext)

	// ExitDeclarationExplicitSlice is called when exiting the DeclarationExplicitSlice production.
	ExitDeclarationExplicitSlice(c *DeclarationExplicitSliceContext)

	// ExitDeclarationImplicitSlice is called when exiting the DeclarationImplicitSlice production.
	ExitDeclarationImplicitSlice(c *DeclarationImplicitSliceContext)

	// ExitStructDeclaration is called when exiting the StructDeclaration production.
	ExitStructDeclaration(c *StructDeclarationContext)

	// ExitStructInstanceDeclaration is called when exiting the StructInstanceDeclaration production.
	ExitStructInstanceDeclaration(c *StructInstanceDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the FunctionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitNormalTypeParameter is called when exiting the NormalTypeParameter production.
	ExitNormalTypeParameter(c *NormalTypeParameterContext)

	// ExitSliceTypeParameter is called when exiting the SliceTypeParameter production.
	ExitSliceTypeParameter(c *SliceTypeParameterContext)

	// ExitTiposSlice is called when exiting the tiposSlice production.
	ExitTiposSlice(c *TiposSliceContext)

	// ExitStruct_field is called when exiting the struct_field production.
	ExitStruct_field(c *Struct_fieldContext)

	// ExitLstStruct is called when exiting the lstStruct production.
	ExitLstStruct(c *LstStructContext)

	// ExitListSlice is called when exiting the listSlice production.
	ExitListSlice(c *ListSliceContext)

	// ExitAssignmentExpr is called when exiting the AssignmentExpr production.
	ExitAssignmentExpr(c *AssignmentExprContext)

	// ExitAssignmentPlusExpr is called when exiting the AssignmentPlusExpr production.
	ExitAssignmentPlusExpr(c *AssignmentPlusExprContext)

	// ExitAssignmentMinusExpr is called when exiting the AssignmentMinusExpr production.
	ExitAssignmentMinusExpr(c *AssignmentMinusExprContext)

	// ExitAssignmentIncrementExpr is called when exiting the AssignmentIncrementExpr production.
	ExitAssignmentIncrementExpr(c *AssignmentIncrementExprContext)

	// ExitAssignmentDecrementExpr is called when exiting the AssignmentDecrementExpr production.
	ExitAssignmentDecrementExpr(c *AssignmentDecrementExprContext)

	// ExitSliceAssignmentExpr is called when exiting the SliceAssignmentExpr production.
	ExitSliceAssignmentExpr(c *SliceAssignmentExprContext)

	// ExitAttributeAssignmentExpr is called when exiting the AttributeAssignmentExpr production.
	ExitAttributeAssignmentExpr(c *AttributeAssignmentExprContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitPrintList is called when exiting the printList production.
	ExitPrintList(c *PrintListContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitElseIfBlock is called when exiting the elseIfBlock production.
	ExitElseIfBlock(c *ElseIfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitSwitch is called when exiting the switch production.
	ExitSwitch(c *SwitchContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitDefaultBlock is called when exiting the defaultBlock production.
	ExitDefaultBlock(c *DefaultBlockContext)

	// ExitForSimple is called when exiting the ForSimple production.
	ExitForSimple(c *ForSimpleContext)

	// ExitForDeclaration is called when exiting the ForDeclaration production.
	ExitForDeclaration(c *ForDeclarationContext)

	// ExitForInSlice is called when exiting the ForInSlice production.
	ExitForInSlice(c *ForInSliceContext)

	// ExitIndividualBlock is called when exiting the individualBlock production.
	ExitIndividualBlock(c *IndividualBlockContext)

	// ExitBreakStatement is called when exiting the BreakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the ContinueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the ReturnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitIndexOfFunction is called when exiting the IndexOfFunction production.
	ExitIndexOfFunction(c *IndexOfFunctionContext)

	// ExitJoinFunction is called when exiting the JoinFunction production.
	ExitJoinFunction(c *JoinFunctionContext)

	// ExitLengthFunction is called when exiting the LengthFunction production.
	ExitLengthFunction(c *LengthFunctionContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitFloatExpr is called when exiting the FloatExpr production.
	ExitFloatExpr(c *FloatExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitDerefExpr is called when exiting the DerefExpr production.
	ExitDerefExpr(c *DerefExprContext)

	// ExitParseFloatExpr is called when exiting the ParseFloatExpr production.
	ExitParseFloatExpr(c *ParseFloatExprContext)

	// ExitNegExpr is called when exiting the NegExpr production.
	ExitNegExpr(c *NegExprContext)

	// ExitAppendExpr is called when exiting the AppendExpr production.
	ExitAppendExpr(c *AppendExprContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitSlicePrimitiveValue is called when exiting the SlicePrimitiveValue production.
	ExitSlicePrimitiveValue(c *SlicePrimitiveValueContext)

	// ExitTypeOfExpr is called when exiting the TypeOfExpr production.
	ExitTypeOfExpr(c *TypeOfExprContext)

	// ExitStructAccess is called when exiting the StructAccess production.
	ExitStructAccess(c *StructAccessContext)

	// ExitFunctionCallExpr is called when exiting the FunctionCallExpr production.
	ExitFunctionCallExpr(c *FunctionCallExprContext)

	// ExitAtoiExpr is called when exiting the AtoiExpr production.
	ExitAtoiExpr(c *AtoiExprContext)

	// ExitParExpr is called when exiting the ParExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitStrExpr is called when exiting the StrExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitRefExpr is called when exiting the RefExpr production.
	ExitRefExpr(c *RefExprContext)

	// ExitSliceAccess is called when exiting the SliceAccess production.
	ExitSliceAccess(c *SliceAccessContext)
}
