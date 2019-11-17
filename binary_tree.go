package datastructs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var EmptyInputErr = errors.New("empty input")
var InvalidInputErr = errors.New("invalid input")

type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
	Value int
}

func NewBinaryTree(vals ...interface{}) (res *BinaryTreeNode, err error) {
	if len(vals) == 0 {
		return nil, EmptyInputErr
	}
	if res, err = createNode(vals[0]); err != nil {
		return
	}
	err = buildTree([]*BinaryTreeNode{res}, 1, vals)
	return
}

func buildTree(nodes []*BinaryTreeNode, level int, vals []interface{}) (err error) {
	count := 1 << level
	offset := count - 1
	if offset+count > len(vals) {
		count = len(vals) - offset
		if count <= 0 {
			return
		}
	}
	nextNodes := make([]*BinaryTreeNode, count)
	var nextNode *BinaryTreeNode
	for i := 0; i < count; i++ {
		if nextNode, err = createNode(vals[i+offset]); err != nil {
			return indexError(i+offset, err)
		}
		if nextNode == nil {
			continue
		}
		node := nodes[i/2]
		if node == nil {
			prevOffset := 1<<(level-1) - 1
			parentIndex := prevOffset + count/2
			return indexError(
				i+offset,
				fmt.Errorf("parent node at index %d is nil: %w", parentIndex, InvalidInputErr),
			)
		}
		if i%2 == 0 {
			node.Left = nextNode
		} else {
			node.Right = nextNode
		}
		nextNodes[i] = nextNode
	}
	return buildTree(nextNodes, level+1, vals)
}

func indexError(index int, err error) error {
	return fmt.Errorf("index %d: %w", index, err)
}

func createNode(val interface{}) (*BinaryTreeNode, error) {
	if val == nil {
		return nil, nil
	}
	if v, ok := val.(int); ok {
		return &BinaryTreeNode{Value: v}, nil
	}
	return nil, fmt.Errorf("unexpected value: %[1]v (%[1]T): %w", val, InvalidInputErr)
}

func (node *BinaryTreeNode) String() string {
	pd := node.getPlotData()
	sb := &strings.Builder{}
	nb := &strings.Builder{}

	node.bft(func(n *BinaryTreeNode, level, index int) {
		d := pd[level-1]
		if level > 1 {
			n.plotEdge(sb, d, index)
		}
		n.plotNode(nb, d, index)
		if index == d.count-1 {
			sb.WriteString(nb.String())
			nb.Reset()
		}
	})
	return sb.String()
}

func (node *BinaryTreeNode) plotEdge(sb *strings.Builder, pd *plotData, index int) {
	var l, m, r string
	innerLineLen := pd.siblingInnerSpacing + (pd.valWidth-pd.valRoofOutset-1-1)*2

	var space, line, left, right = " ", "¯", "/", "\\"
	if node == nil {
		line, left, right = space, space, space
	}

	switch index {
	case 0:
		l = strings.Repeat(space, pd.levelMargin+pd.valRoofOutset+1)
		m = left
		r = strings.Repeat(line, innerLineLen/2+innerLineLen%2)
	case pd.count - 1:
		l = strings.Repeat(line, innerLineLen/2)
		m = right
		r = fmt.Sprintln()
	default:
		isLeft := index%2 == 0
		if isLeft {
			l = ""
			m = left
			r = strings.Repeat(line, innerLineLen/2+innerLineLen%2)
		} else {
			outerSpacing := pd.siblingOuterSpacing + pd.valRoofOutset*2 + 1 + 1
			l = strings.Repeat(line, innerLineLen/2)
			m = right
			r = strings.Repeat(space, outerSpacing)
		}
	}
	sb.WriteString(l)
	sb.WriteString(m)
	sb.WriteString(r)
}

func (node *BinaryTreeNode) plotNode(sb *strings.Builder, pd *plotData, index int) {
	if index == 0 {
		sb.WriteString(strings.Repeat(" ", pd.levelMargin))
	}

	var val string
	if node == nil {
		val = strings.Repeat(" ", pd.valWidth)
	} else {
		var pad string
		var valWidth = pd.valWidth
		if pd.valRoofOutset > 0 {
			if factValWidth := len(strconv.Itoa(node.Value)); factValWidth < pd.valWidth {
				factRoofOutset := pd.valRoofOutset - (pd.valWidth - factValWidth)
				if factRoofOutset < 0 {
					factRoofOutset = 0
				}
				padSize := pd.valRoofOutset - factRoofOutset
				valWidth -= padSize
				pad = strings.Repeat(" ", padSize)
			}
		}
		var padDir, prePad, postPad string
		if index%2 == 0 {
			padDir = "-"
			prePad = pad
		} else {
			postPad = pad
		}
		format := fmt.Sprintf("%s%%%s%dd%s", prePad, padDir, valWidth, postPad)
		val = fmt.Sprintf(format, node.Value)

	}
	sb.WriteString(val)

	if index == pd.count-1 {
		sb.WriteString(fmt.Sprintln())
		return
	}

	var spacing int
	if index%2 == 0 {
		spacing = pd.siblingInnerSpacing
	} else {
		spacing = pd.siblingOuterSpacing
	}
	sb.WriteString(strings.Repeat(" ", spacing))
}

func (node *BinaryTreeNode) bft(cb func(node *BinaryTreeNode, level, index int)) {
	nodes := []*BinaryTreeNode{node}
	level := 1
	run := true
	for run {
		run = false
		nextNodes := make([]*BinaryTreeNode, 0, len(nodes)*2)
		for i, n := range nodes {
			cb(n, level, i)
			if n == nil {
				nextNodes = append(nextNodes, nil, nil)
				continue
			}
			nextNodes = append(nextNodes, n.Left, n.Right)
			if !run {
				run = n.Left != nil || n.Right != nil
			}
		}
		nodes = nextNodes
		level++
	}
}

type plotData struct {
	count               int
	pairWidth           int
	valWidth            int
	levelMargin         int
	valRoofOutset       int
	siblingInnerSpacing int
	siblingOuterSpacing int
}

func (p *plotData) levelUp() *plotData {
	if p.count == 1 {
		return nil
	}
	levelMargin := p.pairWidth/2 - p.valWidth/2
	pairWidth := p.pairWidth*2 + 1
	siblingSpacing := p.pairWidth - p.valWidth + 1
	return &plotData{
		count:               p.count / 2,
		pairWidth:           pairWidth,
		valWidth:            p.valWidth,
		levelMargin:         levelMargin,
		valRoofOutset:       0,
		siblingInnerSpacing: siblingSpacing,
		siblingOuterSpacing: siblingSpacing,
	}
}

func (node *BinaryTreeNode) getPlotData() []*plotData {
	depth, maxVal := node.stat()
	valWidth := len(strconv.Itoa(maxVal))
	siblingInnerSpacing := 2 - valWidth%2 // 1 or 2
	prelimBasePairWidth := valWidth + 4   // 2 obliques + 2 at roof egdes
	factBasePairWidth := valWidth*2 + (2 - valWidth%2)
	if factBasePairWidth < prelimBasePairWidth {
		factBasePairWidth = prelimBasePairWidth
		siblingInnerSpacing = prelimBasePairWidth - valWidth*2
	}
	valRoofOutset := (factBasePairWidth - prelimBasePairWidth) / 2

	levelData := &plotData{
		count:               1 << (depth - 1),
		pairWidth:           factBasePairWidth,
		valWidth:            valWidth,
		levelMargin:         0,
		valRoofOutset:       valRoofOutset,
		siblingInnerSpacing: siblingInnerSpacing,
		siblingOuterSpacing: 1,
	}
	res := make([]*plotData, depth)
	for i := depth - 1; i >= 0; i-- {
		res[i] = levelData
		levelData = levelData.levelUp()
	}
	return res
}

func (node *BinaryTreeNode) stat() (depth, maxVal int) {
	depth, maxVal = node.statRec(0, MinInt)
	return
}

func (node *BinaryTreeNode) statRec(depth, val int) (int, int) {
	if node == nil {
		return depth, val
	}
	depth++
	if node.Value > val {
		val = node.Value
	}
	ld, lv := node.Left.statRec(depth, val)
	rd, rv := node.Right.statRec(depth, val)
	if ld > rd {
		depth = ld
	} else {
		depth = rd
	}
	if lv > rv {
		if lv > val {
			val = lv
		}
	} else if rv > val {
		val = rv
	}
	return depth, val
}
