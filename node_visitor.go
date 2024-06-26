package inspect

import (
	"github.com/ZenLiuCN/fn"
	"go/ast"
)

type (
	Nodes = fn.Stack[ast.Node]

	NodeDeclVisit[I Inspector, T ast.Node, X any] func(i I, d Dir, name string, o *ast.Object, x T, seen Nodes, cx X) bool
	NodeVisit[I Inspector, T ast.Node, X any]     func(i I, d Dir, o *ast.Object, x T, seen Nodes, cx X) bool
	NodeExprVisit[I Inspector, T ast.Expr, X any] func(i I, d Dir, o *ast.Object, x T, seen Nodes, cx X) bool
	NodeVisitor[I Inspector, X any]               interface {
		VisitFieldDecl(i I, d Dir, name string, o *ast.Object, x *ast.Field, seen Nodes, cx X) bool
		VisitFuncDecl(i I, d Dir, name string, o *ast.Object, x *ast.FuncDecl, nodes Nodes, cx X) bool
		VisitTypeDecl(i I, d Dir, name string, o *ast.Object, x *ast.TypeSpec, nodes Nodes, cx X) bool
		VisitValueDecl(i I, d Dir, name string, o *ast.Object, x *ast.ValueSpec, nodes Nodes, cx X) bool

		VisitFieldList(i I, d Dir, o *ast.Object, x *ast.FieldList, nodes Nodes, cx X) bool
		VisitField(i I, d Dir, o *ast.Object, x *ast.Field, seen Nodes, cx X) bool

		VisitIdent(i I, d Dir, o *ast.Object, x *ast.Ident, nodes Nodes, cx X) bool
		VisitFuncType(i I, d Dir, o *ast.Object, x *ast.FuncType, nodes Nodes, cx X) bool
		VisitStructType(i I, d Dir, o *ast.Object, x *ast.StructType, nodes Nodes, cx X) bool
		VisitInterfaceType(i I, d Dir, o *ast.Object, x *ast.InterfaceType, nodes Nodes, cx X) bool
		VisitArrayType(i I, d Dir, o *ast.Object, x *ast.ArrayType, nodes Nodes, cx X) bool
		VisitMapType(i I, d Dir, o *ast.Object, x *ast.MapType, nodes Nodes, cx X) bool
		VisitChanType(i I, d Dir, o *ast.Object, x *ast.ChanType, nodes Nodes, cx X) bool
		VisitStarExpr(i I, d Dir, o *ast.Object, x *ast.StarExpr, nodes Nodes, cx X) bool
		VisitSelectorExpr(i I, d Dir, o *ast.Object, x *ast.SelectorExpr, nodes Nodes, cx X) bool
		VisitBasicLit(i I, d Dir, o *ast.Object, x *ast.BasicLit, nodes Nodes, cx X) bool
		VisitEllipsis(i I, d Dir, o *ast.Object, x *ast.Ellipsis, nodes Nodes, cx X) bool
		VisitUnaryExpr(i I, d Dir, o *ast.Object, x *ast.UnaryExpr, nodes Nodes, cx X) bool
		VisitBinaryExpr(i I, d Dir, o *ast.Object, x *ast.BinaryExpr, nodes Nodes, cx X) bool
	}
	FnNodeVisitor[I Inspector, X any] struct {
		FnVisitFieldDecl NodeDeclVisit[I, *ast.Field, X]
		FnVisitFuncDecl  NodeDeclVisit[I, *ast.FuncDecl, X]
		FnVisitTypeDecl  NodeDeclVisit[I, *ast.TypeSpec, X]
		FnVisitValueDecl NodeDeclVisit[I, *ast.ValueSpec, X]

		FnVisitFieldList NodeVisit[I, *ast.FieldList, X]
		FnVisitField     NodeVisit[I, *ast.Field, X]

		FnVisitIdent         NodeExprVisit[I, *ast.Ident, X]
		FnVisitFuncType      NodeExprVisit[I, *ast.FuncType, X]
		FnVisitStructType    NodeExprVisit[I, *ast.StructType, X]
		FnVisitInterfaceType NodeExprVisit[I, *ast.InterfaceType, X]
		FnVisitArrayType     NodeExprVisit[I, *ast.ArrayType, X]
		FnVisitMapType       NodeExprVisit[I, *ast.MapType, X]
		FnVisitChanType      NodeExprVisit[I, *ast.ChanType, X]
		FnVisitStarExpr      NodeExprVisit[I, *ast.StarExpr, X]
		FnVisitSelectorExpr  NodeExprVisit[I, *ast.SelectorExpr, X]
		FnVisitBasicLit      NodeExprVisit[I, *ast.BasicLit, X]
		FnVisitEllipsis      NodeExprVisit[I, *ast.Ellipsis, X]
		FnVisitUnaryExpr     NodeExprVisit[I, *ast.UnaryExpr, X]
		FnVisitBinaryExpr    NodeExprVisit[I, *ast.BinaryExpr, X]
	}
)

func (n FnNodeVisitor[I, X]) VisitFieldDecl(i I, d Dir, name string, o *ast.Object, x *ast.Field, nodes Nodes, cx X) bool {
	if n.FnVisitFieldDecl != nil {
		return n.FnVisitFieldDecl(i, d, name, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitFuncDecl(i I, d Dir, name string, o *ast.Object, x *ast.FuncDecl, nodes Nodes, cx X) bool {
	if n.FnVisitFuncDecl != nil {
		return n.FnVisitFuncDecl(i, d, name, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitTypeDecl(i I, d Dir, name string, o *ast.Object, x *ast.TypeSpec, nodes Nodes, cx X) bool {
	if n.FnVisitTypeDecl != nil {
		return n.FnVisitTypeDecl(i, d, name, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitValueDecl(i I, d Dir, name string, o *ast.Object, x *ast.ValueSpec, nodes Nodes, cx X) bool {
	if n.FnVisitValueDecl != nil {
		return n.FnVisitValueDecl(i, d, name, o, x, nodes, cx)
	}
	return true
}

func (n FnNodeVisitor[I, X]) VisitFieldList(i I, d Dir, o *ast.Object, x *ast.FieldList, nodes Nodes, cx X) bool {
	if n.FnVisitFieldList != nil {
		return n.FnVisitFieldList(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitField(i I, d Dir, o *ast.Object, x *ast.Field, nodes Nodes, cx X) bool {
	if n.FnVisitField != nil {
		return n.FnVisitField(i, d, o, x, nodes, cx)
	}
	return true
}

func (n FnNodeVisitor[I, X]) VisitIdent(i I, d Dir, o *ast.Object, x *ast.Ident, nodes Nodes, cx X) bool {
	if n.FnVisitIdent != nil {
		return n.FnVisitIdent(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitFuncType(i I, d Dir, o *ast.Object, x *ast.FuncType, nodes Nodes, cx X) bool {
	if n.FnVisitFuncType != nil {
		return n.FnVisitFuncType(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitStructType(i I, d Dir, o *ast.Object, x *ast.StructType, nodes Nodes, cx X) bool {
	if n.FnVisitStructType != nil {
		return n.FnVisitStructType(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitInterfaceType(i I, d Dir, o *ast.Object, x *ast.InterfaceType, nodes Nodes, cx X) bool {
	if n.FnVisitInterfaceType != nil {
		return n.FnVisitInterfaceType(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitArrayType(i I, d Dir, o *ast.Object, x *ast.ArrayType, nodes Nodes, cx X) bool {
	if n.FnVisitArrayType != nil {
		return n.FnVisitArrayType(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitMapType(i I, d Dir, o *ast.Object, x *ast.MapType, nodes Nodes, cx X) bool {
	if n.FnVisitMapType != nil {
		return n.FnVisitMapType(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitChanType(i I, d Dir, o *ast.Object, x *ast.ChanType, nodes Nodes, cx X) bool {
	if n.FnVisitChanType != nil {
		return n.FnVisitChanType(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitStarExpr(i I, d Dir, o *ast.Object, x *ast.StarExpr, nodes Nodes, cx X) bool {
	if n.FnVisitStarExpr != nil {
		return n.FnVisitStarExpr(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitSelectorExpr(i I, d Dir, o *ast.Object, x *ast.SelectorExpr, nodes Nodes, cx X) bool {
	if n.FnVisitSelectorExpr != nil {
		return n.FnVisitSelectorExpr(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitBasicLit(i I, d Dir, o *ast.Object, x *ast.BasicLit, nodes Nodes, cx X) bool {
	if n.FnVisitBasicLit != nil {
		return n.FnVisitBasicLit(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitEllipsis(i I, d Dir, o *ast.Object, x *ast.Ellipsis, nodes Nodes, cx X) bool {
	if n.FnVisitEllipsis != nil {
		return n.FnVisitEllipsis(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitUnaryExpr(i I, d Dir, o *ast.Object, x *ast.UnaryExpr, nodes Nodes, cx X) bool {
	if n.FnVisitUnaryExpr != nil {
		return n.FnVisitUnaryExpr(i, d, o, x, nodes, cx)
	}
	return true
}
func (n FnNodeVisitor[I, X]) VisitBinaryExpr(i I, d Dir, o *ast.Object, x *ast.BinaryExpr, nodes Nodes, cx X) bool {
	if n.FnVisitBinaryExpr != nil {
		return n.FnVisitBinaryExpr(i, d, o, x, nodes, cx)
	}
	return true
}
func IsNode[t ast.Node](expr ast.Node) (ok bool) {
	_, ok = expr.(t)
	return
}
func IsNodeExpr[t ast.Expr](expr ast.Expr) (ok bool) {
	_, ok = expr.(t)
	return
}

var (
	IsNodeExprIdent          = IsNodeExpr[*ast.Ident]
	IsNodeExprFuncType       = IsNodeExpr[*ast.FuncType]
	IsNodeExprStructType     = IsNodeExpr[*ast.StructType]
	IsNodeExprInterfaceType  = IsNodeExpr[*ast.InterfaceType]
	IsNodeExprArrayType      = IsNodeExpr[*ast.ArrayType]
	IsNodeExprMapType        = IsNodeExpr[*ast.MapType]
	IsNodeExprChanType       = IsNodeExpr[*ast.ChanType]
	IsNodeExprStarExpr       = IsNodeExpr[*ast.StarExpr]
	IsNodeExprSelectorExpr   = IsNodeExpr[*ast.SelectorExpr]
	IsNodeExprBasicLit       = IsNodeExpr[*ast.BasicLit]
	IsNodeExprEllipsis       = IsNodeExpr[*ast.Ellipsis]
	IsNodeExprUnaryExpr      = IsNodeExpr[*ast.UnaryExpr]
	IsNodeExprBinaryExpr     = IsNodeExpr[*ast.BinaryExpr]
	IsNodeExprSliceExpr      = IsNodeExpr[*ast.SliceExpr]
	IsNodeExprIndexExpr      = IsNodeExpr[*ast.IndexExpr]
	IsNodeExprCallExpr       = IsNodeExpr[*ast.CallExpr]
	IsNodeExprTypeAssertExpr = IsNodeExpr[*ast.TypeAssertExpr]
	IsNodeExprKeyValueExpr   = IsNodeExpr[*ast.KeyValueExpr]
	IsNodeExprBadExpr        = IsNodeExpr[*ast.BadExpr]
	IsNodeExprCompositeLit   = IsNodeExpr[*ast.CompositeLit]
	IsNodeExprFuncLit        = IsNodeExpr[*ast.FuncLit]
	IsNodeExprParenExpr      = IsNodeExpr[*ast.ParenExpr]
)

var (
	IsNodeField       = IsNode[*ast.Field]
	IsNodeValueSpec   = IsNode[*ast.ValueSpec]
	IsNodeTypeSpec    = IsNode[*ast.TypeSpec]
	IsNodeFuncDecl    = IsNode[*ast.FuncDecl]
	IsNodeLabeledStmt = IsNode[*ast.LabeledStmt]
	IsNodeAssignStmt  = IsNode[*ast.AssignStmt]
	//IsNodeScope       = IsNode[*ast.Scope]

	IsNodeFieldList = IsNode[*ast.FieldList]

	IsNodeIdent          = IsNode[*ast.Ident]
	IsNodeFuncType       = IsNode[*ast.FuncType]
	IsNodeStructType     = IsNode[*ast.StructType]
	IsNodeInterfaceType  = IsNode[*ast.InterfaceType]
	IsNodeArrayType      = IsNode[*ast.ArrayType]
	IsNodeMapType        = IsNode[*ast.MapType]
	IsNodeChanType       = IsNode[*ast.ChanType]
	IsNodeStarExpr       = IsNode[*ast.StarExpr]
	IsNodeSelectorExpr   = IsNode[*ast.SelectorExpr]
	IsNodeBasicLit       = IsNode[*ast.BasicLit]
	IsNodeEllipsis       = IsNode[*ast.Ellipsis]
	IsNodeUnaryExpr      = IsNode[*ast.UnaryExpr]
	IsNodeBinaryExpr     = IsNode[*ast.BinaryExpr]
	IsNodeSliceExpr      = IsNode[*ast.SliceExpr]
	IsNodeIndexExpr      = IsNode[*ast.IndexExpr]
	IsNodeCallExpr       = IsNode[*ast.CallExpr]
	IsNodeTypeAssertExpr = IsNode[*ast.TypeAssertExpr]
	IsNodeKeyValueExpr   = IsNode[*ast.KeyValueExpr]
	IsNodeBadExpr        = IsNode[*ast.BadExpr]
	IsNodeCompositeLit   = IsNode[*ast.CompositeLit]
	IsNodeFuncLit        = IsNode[*ast.FuncLit]
	IsNodeParenExpr      = IsNode[*ast.ParenExpr]
)
