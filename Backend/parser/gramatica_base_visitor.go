// Code generated from Gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr4-go/antlr/v4"

type BaseGramaticaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGramaticaVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitMainfunction(ctx *MainfunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitDeclarationExplicitVar(ctx *DeclarationExplicitVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitDeclarationImplicitVar(ctx *DeclarationImplicitVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitDeclarationExplicitSlice(ctx *DeclarationExplicitSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitDeclarationImplicitSlice(ctx *DeclarationImplicitSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitStructInstanceDeclaration(ctx *StructInstanceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitNormalTypeParameter(ctx *NormalTypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitSliceTypeParameter(ctx *SliceTypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitTiposSlice(ctx *TiposSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitStruct_field(ctx *Struct_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitLstStruct(ctx *LstStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitListSlice(ctx *ListSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitAssignmentExpr(ctx *AssignmentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitAssignmentPlusExpr(ctx *AssignmentPlusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitAssignmentMinusExpr(ctx *AssignmentMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitAssignmentIncrementExpr(ctx *AssignmentIncrementExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitAssignmentDecrementExpr(ctx *AssignmentDecrementExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitSliceAssignmentExpr(ctx *SliceAssignmentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitAttributeAssignmentExpr(ctx *AttributeAssignmentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitPrint(ctx *PrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitPrintList(ctx *PrintListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitIf(ctx *IfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitElseIfBlock(ctx *ElseIfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitElseBlock(ctx *ElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitSwitch(ctx *SwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitDefaultBlock(ctx *DefaultBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitForSimple(ctx *ForSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitForDeclaration(ctx *ForDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitForInSlice(ctx *ForInSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitIndividualBlock(ctx *IndividualBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitIndexOfFunction(ctx *IndexOfFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitJoinFunction(ctx *JoinFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitLengthFunction(ctx *LengthFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitDerefExpr(ctx *DerefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitParseFloatExpr(ctx *ParseFloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitNegExpr(ctx *NegExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitAppendExpr(ctx *AppendExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitOpExpr(ctx *OpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitSlicePrimitiveValue(ctx *SlicePrimitiveValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitTypeOfExpr(ctx *TypeOfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitStructAccess(ctx *StructAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitAtoiExpr(ctx *AtoiExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitStrExpr(ctx *StrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitRefExpr(ctx *RefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitSliceAccess(ctx *SliceAccessContext) interface{} {
	return v.VisitChildren(ctx)
}
