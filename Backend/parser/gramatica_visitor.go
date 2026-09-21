// Code generated from Gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GramaticaParser.
type GramaticaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GramaticaParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GramaticaParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GramaticaParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by GramaticaParser#mainfunction.
	VisitMainfunction(ctx *MainfunctionContext) interface{}

	// Visit a parse tree produced by GramaticaParser#DeclarationExplicitVar.
	VisitDeclarationExplicitVar(ctx *DeclarationExplicitVarContext) interface{}

	// Visit a parse tree produced by GramaticaParser#DeclarationImplicitVar.
	VisitDeclarationImplicitVar(ctx *DeclarationImplicitVarContext) interface{}

	// Visit a parse tree produced by GramaticaParser#DeclarationExplicitSlice.
	VisitDeclarationExplicitSlice(ctx *DeclarationExplicitSliceContext) interface{}

	// Visit a parse tree produced by GramaticaParser#DeclarationImplicitSlice.
	VisitDeclarationImplicitSlice(ctx *DeclarationImplicitSliceContext) interface{}

	// Visit a parse tree produced by GramaticaParser#StructDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by GramaticaParser#StructInstanceDeclaration.
	VisitStructInstanceDeclaration(ctx *StructInstanceDeclarationContext) interface{}

	// Visit a parse tree produced by GramaticaParser#FunctionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by GramaticaParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by GramaticaParser#NormalTypeParameter.
	VisitNormalTypeParameter(ctx *NormalTypeParameterContext) interface{}

	// Visit a parse tree produced by GramaticaParser#SliceTypeParameter.
	VisitSliceTypeParameter(ctx *SliceTypeParameterContext) interface{}

	// Visit a parse tree produced by GramaticaParser#tiposSlice.
	VisitTiposSlice(ctx *TiposSliceContext) interface{}

	// Visit a parse tree produced by GramaticaParser#struct_field.
	VisitStruct_field(ctx *Struct_fieldContext) interface{}

	// Visit a parse tree produced by GramaticaParser#lstStruct.
	VisitLstStruct(ctx *LstStructContext) interface{}

	// Visit a parse tree produced by GramaticaParser#listSlice.
	VisitListSlice(ctx *ListSliceContext) interface{}

	// Visit a parse tree produced by GramaticaParser#AssignmentExpr.
	VisitAssignmentExpr(ctx *AssignmentExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#AssignmentPlusExpr.
	VisitAssignmentPlusExpr(ctx *AssignmentPlusExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#AssignmentMinusExpr.
	VisitAssignmentMinusExpr(ctx *AssignmentMinusExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#AssignmentIncrementExpr.
	VisitAssignmentIncrementExpr(ctx *AssignmentIncrementExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#AssignmentDecrementExpr.
	VisitAssignmentDecrementExpr(ctx *AssignmentDecrementExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#SliceAssignmentExpr.
	VisitSliceAssignmentExpr(ctx *SliceAssignmentExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#AttributeAssignmentExpr.
	VisitAttributeAssignmentExpr(ctx *AttributeAssignmentExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by GramaticaParser#printList.
	VisitPrintList(ctx *PrintListContext) interface{}

	// Visit a parse tree produced by GramaticaParser#if.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by GramaticaParser#elseIfBlock.
	VisitElseIfBlock(ctx *ElseIfBlockContext) interface{}

	// Visit a parse tree produced by GramaticaParser#elseBlock.
	VisitElseBlock(ctx *ElseBlockContext) interface{}

	// Visit a parse tree produced by GramaticaParser#switch.
	VisitSwitch(ctx *SwitchContext) interface{}

	// Visit a parse tree produced by GramaticaParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by GramaticaParser#defaultBlock.
	VisitDefaultBlock(ctx *DefaultBlockContext) interface{}

	// Visit a parse tree produced by GramaticaParser#ForSimple.
	VisitForSimple(ctx *ForSimpleContext) interface{}

	// Visit a parse tree produced by GramaticaParser#ForDeclaration.
	VisitForDeclaration(ctx *ForDeclarationContext) interface{}

	// Visit a parse tree produced by GramaticaParser#ForInSlice.
	VisitForInSlice(ctx *ForInSliceContext) interface{}

	// Visit a parse tree produced by GramaticaParser#individualBlock.
	VisitIndividualBlock(ctx *IndividualBlockContext) interface{}

	// Visit a parse tree produced by GramaticaParser#BreakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by GramaticaParser#ContinueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by GramaticaParser#ReturnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by GramaticaParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by GramaticaParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by GramaticaParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by GramaticaParser#IndexOfFunction.
	VisitIndexOfFunction(ctx *IndexOfFunctionContext) interface{}

	// Visit a parse tree produced by GramaticaParser#JoinFunction.
	VisitJoinFunction(ctx *JoinFunctionContext) interface{}

	// Visit a parse tree produced by GramaticaParser#LengthFunction.
	VisitLengthFunction(ctx *LengthFunctionContext) interface{}

	// Visit a parse tree produced by GramaticaParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#FloatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#DerefExpr.
	VisitDerefExpr(ctx *DerefExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#ParseFloatExpr.
	VisitParseFloatExpr(ctx *ParseFloatExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#NegExpr.
	VisitNegExpr(ctx *NegExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#AppendExpr.
	VisitAppendExpr(ctx *AppendExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#SlicePrimitiveValue.
	VisitSlicePrimitiveValue(ctx *SlicePrimitiveValueContext) interface{}

	// Visit a parse tree produced by GramaticaParser#TypeOfExpr.
	VisitTypeOfExpr(ctx *TypeOfExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#StructAccess.
	VisitStructAccess(ctx *StructAccessContext) interface{}

	// Visit a parse tree produced by GramaticaParser#FunctionCallExpr.
	VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#AtoiExpr.
	VisitAtoiExpr(ctx *AtoiExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#RefExpr.
	VisitRefExpr(ctx *RefExprContext) interface{}

	// Visit a parse tree produced by GramaticaParser#SliceAccess.
	VisitSliceAccess(ctx *SliceAccessContext) interface{}
}
