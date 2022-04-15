package examples

// 自定义的包2
import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/minnanalee/antlrjava/parser"
)

type BaseJavaParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJavaParserVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitPackageDeclaration(ctx *parser.PackageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitImportDeclaration(ctx *parser.ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitModifier(ctx *parser.ModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitClassOrInterfaceModifier(ctx *parser.ClassOrInterfaceModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitVariableModifier(ctx *parser.VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeParameters(ctx *parser.TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeParameter(ctx *parser.TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeBound(ctx *parser.TypeBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitEnumConstants(ctx *parser.EnumConstantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitEnumConstant(ctx *parser.EnumConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitEnumBodyDeclarations(ctx *parser.EnumBodyDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitInterfaceBody(ctx *parser.InterfaceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitMethodBody(ctx *parser.MethodBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeTypeOrVoid(ctx *parser.TypeTypeOrVoidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitGenericMethodDeclaration(ctx *parser.GenericMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitGenericConstructorDeclaration(ctx *parser.GenericConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitInterfaceMemberDeclaration(ctx *parser.InterfaceMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitConstDeclaration(ctx *parser.ConstDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitConstantDeclarator(ctx *parser.ConstantDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitInterfaceMethodModifier(ctx *parser.InterfaceMethodModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitGenericInterfaceMethodDeclaration(ctx *parser.GenericInterfaceMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitInterfaceCommonBodyDeclaration(ctx *parser.InterfaceCommonBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitVariableInitializer(ctx *parser.VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitArrayInitializer(ctx *parser.ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitClassOrInterfaceType(ctx *parser.ClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeArgument(ctx *parser.TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitQualifiedNameList(ctx *parser.QualifiedNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitFormalParameters(ctx *parser.FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitReceiverParameter(ctx *parser.ReceiverParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitFormalParameterList(ctx *parser.FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitFormalParameter(ctx *parser.FormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLastFormalParameter(ctx *parser.LastFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLambdaLVTIList(ctx *parser.LambdaLVTIListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLambdaLVTIParameter(ctx *parser.LambdaLVTIParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitQualifiedName(ctx *parser.QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitIntegerLiteral(ctx *parser.IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAltAnnotationQualifiedName(ctx *parser.AltAnnotationQualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAnnotation(ctx *parser.AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitElementValuePairs(ctx *parser.ElementValuePairsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitElementValuePair(ctx *parser.ElementValuePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitElementValue(ctx *parser.ElementValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitElementValueArrayInitializer(ctx *parser.ElementValueArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAnnotationTypeDeclaration(ctx *parser.AnnotationTypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAnnotationTypeBody(ctx *parser.AnnotationTypeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAnnotationTypeElementDeclaration(ctx *parser.AnnotationTypeElementDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAnnotationTypeElementRest(ctx *parser.AnnotationTypeElementRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAnnotationMethodOrConstantRest(ctx *parser.AnnotationMethodOrConstantRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAnnotationMethodRest(ctx *parser.AnnotationMethodRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitAnnotationConstantRest(ctx *parser.AnnotationConstantRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitDefaultValue(ctx *parser.DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitModuleDeclaration(ctx *parser.ModuleDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitModuleBody(ctx *parser.ModuleBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitModuleDirective(ctx *parser.ModuleDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitRequiresModifier(ctx *parser.RequiresModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitRecordDeclaration(ctx *parser.RecordDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitRecordHeader(ctx *parser.RecordHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitRecordComponentList(ctx *parser.RecordComponentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitRecordComponent(ctx *parser.RecordComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitRecordBody(ctx *parser.RecordBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLocalVariableDeclaration(ctx *parser.LocalVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLocalTypeDeclaration(ctx *parser.LocalTypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitCatchClause(ctx *parser.CatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitCatchType(ctx *parser.CatchTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitFinallyBlock(ctx *parser.FinallyBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitResourceSpecification(ctx *parser.ResourceSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitResources(ctx *parser.ResourcesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitResource(ctx *parser.ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitSwitchBlockStatementGroup(ctx *parser.SwitchBlockStatementGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitSwitchLabel(ctx *parser.SwitchLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitForControl(ctx *parser.ForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitForInit(ctx *parser.ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitEnhancedForControl(ctx *parser.EnhancedForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitMethodCall(ctx *parser.MethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitPattern(ctx *parser.PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLambdaExpression(ctx *parser.LambdaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLambdaParameters(ctx *parser.LambdaParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitLambdaBody(ctx *parser.LambdaBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitSwitchExpression(ctx *parser.SwitchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitSwitchLabeledRule(ctx *parser.SwitchLabeledRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitGuardedPattern(ctx *parser.GuardedPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitSwitchRuleOutcome(ctx *parser.SwitchRuleOutcomeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitClassType(ctx *parser.ClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitCreator(ctx *parser.CreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitCreatedName(ctx *parser.CreatedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitInnerCreator(ctx *parser.InnerCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitArrayCreatorRest(ctx *parser.ArrayCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitClassCreatorRest(ctx *parser.ClassCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitExplicitGenericInvocation(ctx *parser.ExplicitGenericInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeArgumentsOrDiamond(ctx *parser.TypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitNonWildcardTypeArgumentsOrDiamond(ctx *parser.NonWildcardTypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitNonWildcardTypeArguments(ctx *parser.NonWildcardTypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeType(ctx *parser.TypeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitPrimitiveType(ctx *parser.PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitSuperSuffix(ctx *parser.SuperSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitExplicitGenericInvocationSuffix(ctx *parser.ExplicitGenericInvocationSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}
