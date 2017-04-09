//Copyright 2017 GoGraphviz Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http)://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package gographviz

import "fmt"

// Attr is an attribute key
type Attr string

// NewAttr creates a new attribute key by checking whether it is a valid key
func NewAttr(key string) (Attr, error) {
	a, ok := validAttrs[key]
	if !ok {
		return Attr(""), fmt.Errorf("%s is not a valid attribute", key)
	}
	return a, nil
}

const (
	Damping            Attr = "Damping"
	K                  Attr = "K"
	URL                Attr = "URL"
	Background         Attr = "_background"
	Area               Attr = "area"
	ArrowHead          Attr = "arrowhead"
	ArrowSize          Attr = "arrowsize"
	ArrowTail          Attr = "arrowtail"
	BB                 Attr = "bb"
	BgColor            Attr = "bgcolor"
	Center             Attr = "center"
	Charset            Attr = "charset"
	ClusterRank        Attr = "clusterrank"
	Color              Attr = "color"
	ColorScheme        Attr = "colorscheme"
	Comment            Attr = "comment"
	Compound           Attr = "compound"
	Concentrate        Attr = "concentrate"
	Constraint         Attr = "constraint"
	Decorate           Attr = "decorate"
	DefaultDist        Attr = "defaultdist"
	Dim                Attr = "dim"
	Dimen              Attr = "dimen"
	Dir                Attr = "dir"
	DirEdgeConstraints Attr = "diredgeconstraints"
	Distortion         Attr = "distortion"
	DPI                Attr = "dpi"
	EdgeURL            Attr = "edgeURL"
	EdgeHREF           Attr = "edgehref"
	EdgeTarget         Attr = "edgetarget"
	EdgeTooltip        Attr = "edgetooltip"
	Epsilon            Attr = "epsilon"
	ESep               Attr = "esep"
	FillColor          Attr = "fillcolor"
	FixedSize          Attr = "fixedsize"
	FontColor          Attr = "fontcolor"
	FontName           Attr = "fontname"
	FontNames          Attr = "fontnames"
	FontPath           Attr = "fontpath"
	FontSize           Attr = "fontsize"
	ForceLabels        Attr = "forcelabels"
	GradientAngle      Attr = "gradientangle"
	Group              Attr = "group"
	HeadURL            Attr = "headURL"
	HeadLP             Attr = "head_lp"
	HeadClip           Attr = "headclip"
	HeadHREF           Attr = "headhref"
	HeadLabel          Attr = "headlabel"
	HeadPort           Attr = "headport"
	HeadTarget         Attr = "headtarget"
	HeadTooltip        Attr = "headtooltip"
	Height             Attr = "height"
	HREF               Attr = "href"
	ID                 Attr = "id"
	Image              Attr = "image"
	ImagePath          Attr = "imagepath"
	ImageScale         Attr = "imagescale"
	InputScale         Attr = "inputscale"
	Label              Attr = "label"
	LabelURL           Attr = "labelURL"
	LabelScheme        Attr = "label_scheme"
	LabelAngle         Attr = "labelangle"
	LabelDistance      Attr = "labeldistance"
	LabelFloat         Attr = "labelfloat"
	LabelFontColor     Attr = "labelfontcolor"
	LabelFontName      Attr = "labelfontname"
	LabelFontSize      Attr = "labelfontsize"
	LabelHREF          Attr = "labelhref"
	LabelJust          Attr = "labeljust"
	LabelLOC           Attr = "labelloc"
	LabelTarget        Attr = "labeltarget"
	LabelTooltip       Attr = "labeltooltip"
	Landscape          Attr = "landscape"
	Layer              Attr = "layer"
	LayerListSep       Attr = "layerlistsep"
	Layers             Attr = "layers"
	LayerSelect        Attr = "layerselect"
	LayerSep           Attr = "layersep"
	Layout             Attr = "layout"
	Len                Attr = "len"
	Levels             Attr = "levels"
	LevelsGap          Attr = "levelsgap"
	LHead              Attr = "lhead"
	LHeight            Attr = "lheight"
	LP                 Attr = "lp"
	LTail              Attr = "ltail"
	LWidth             Attr = "lwidth"
	Margin             Attr = "margin"
	MaxIter            Attr = "maxiter"
	MCLimit            Attr = "mclimit"
	MinDist            Attr = "mindist"
	MinLen             Attr = "minlen"
	Mode               Attr = "mode"
	Model              Attr = "model"
	Mosek              Attr = "mosek"
	NodeSep            Attr = "nodesep"
	NoJustify          Attr = "nojustify"
	Normalize          Attr = "normalize"
	NoTranslate        Attr = "notranslate"
	NSLimit            Attr = "nslimit"
	NSLimit1           Attr = "nslimit1"
	Ordering           Attr = "ordering"
	Orientation        Attr = "orientation"
	Outputorder        Attr = "outputorder"
	Overlap            Attr = "overlap"
	OverlapScaling     Attr = "overlap_scaling"
	OverlapShrink      Attr = "overlap_shrink"
	Pack               Attr = "pack"
	PackMode           Attr = "packmode"
	Pad                Attr = "pad"
	Page               Attr = "page"
	PageDir            Attr = "pagedir"
	PenColor           Attr = "pencolor"
	PenWidth           Attr = "penwidth"
	Peripheries        Attr = "peripheries"
	Pin                Attr = "pin"
	Pos                Attr = "pos"
	Quadtree           Attr = "quadtree"
	Quantum            Attr = "quantum"
	Rank               Attr = "rank"
	RankDir            Attr = "rankdir"
	RankSep            Attr = "ranksep"
	Ratio              Attr = "ratio"
	Rects              Attr = "rects"
	Regular            Attr = "regular"
	ReMinCross         Attr = "remincross"
	RepulsiveForce     Attr = "repulsiveforce"
	Resolution         Attr = "resolution"
	Root               Attr = "root"
	Rotate             Attr = "rotate"
	Rotation           Attr = "rotation"
	SameHead           Attr = "samehead"
	SameTail           Attr = "sametail"
	SamplePoints       Attr = "samplepoints"
	Scale              Attr = "scale"
	SearchSize         Attr = "searchsize"
	Sep                Attr = "sep"
	Shape              Attr = "shape"
	ShapeFile          Attr = "shapefile"
	ShowBoxes          Attr = "showboxes"
	Sides              Attr = "sides"
	Size               Attr = "size"
	Skew               Attr = "skew"
	Smoothing          Attr = "smoothing"
	SortV              Attr = "sortv"
	Splines            Attr = "splines"
	Start              Attr = "start"
	Style              Attr = "style"
	StyleSheet         Attr = "stylesheet"
	TailURL            Attr = "tailURL"
	TailLP             Attr = "tail_lp"
	TailClip           Attr = "tailclip"
	TailHREF           Attr = "tailhref"
	TailLabel          Attr = "taillabel"
	TailPort           Attr = "tailport"
	TailTarget         Attr = "tailtarget"
	TailTooltip        Attr = "tailtooltip"
	Target             Attr = "target"
	Tooltip            Attr = "tooltip"
	TrueColor          Attr = "truecolor"
	Vertices           Attr = "vertices"
	ViewPort           Attr = "viewport"
	VoroMargin         Attr = "voro_margin"
	Weight             Attr = "weight"
	Width              Attr = "width"
	XDotVersion        Attr = "xdotversion"
	XLabel             Attr = "xlabel"
	XLP                Attr = "xlp"
	Z                  Attr = "z"

	MinCross Attr = "mincross" // not in the documentation, but found in the Ped_Lion_Share (lion_share.gv.txt) example
	SSize    Attr = "ssize"    // not in the documentation, but found in the siblings.gv.txt example
	Outline  Attr = "outline"  // not in the documentation, but found in the siblings.gv.txt example
	F        Attr = "f"        // not in the documentation, but found in the transparency.gv.txt example
)

var validAttrs = map[string]Attr{
	string(Damping):            Damping,
	string(K):                  K,
	string(URL):                URL,
	string(Background):         Background,
	string(Area):               Area,
	string(ArrowHead):          ArrowHead,
	string(ArrowSize):          ArrowSize,
	string(ArrowTail):          ArrowTail,
	string(BB):                 BB,
	string(BgColor):            BgColor,
	string(Center):             Center,
	string(Charset):            Charset,
	string(ClusterRank):        ClusterRank,
	string(Color):              Color,
	string(ColorScheme):        ColorScheme,
	string(Comment):            Comment,
	string(Compound):           Compound,
	string(Concentrate):        Concentrate,
	string(Constraint):         Constraint,
	string(Decorate):           Decorate,
	string(DefaultDist):        DefaultDist,
	string(Dim):                Dim,
	string(Dimen):              Dimen,
	string(Dir):                Dir,
	string(DirEdgeConstraints): DirEdgeConstraints,
	string(Distortion):         Distortion,
	string(DPI):                DPI,
	string(EdgeURL):            EdgeURL,
	string(EdgeHREF):           EdgeHREF,
	string(EdgeTarget):         EdgeTarget,
	string(EdgeTooltip):        EdgeTooltip,
	string(Epsilon):            Epsilon,
	string(ESep):               ESep,
	string(FillColor):          FillColor,
	string(FixedSize):          FixedSize,
	string(FontColor):          FontColor,
	string(FontName):           FontName,
	string(FontNames):          FontNames,
	string(FontPath):           FontPath,
	string(FontSize):           FontSize,
	string(ForceLabels):        ForceLabels,
	string(GradientAngle):      GradientAngle,
	string(Group):              Group,
	string(HeadURL):            HeadURL,
	string(HeadLP):             HeadLP,
	string(HeadClip):           HeadClip,
	string(HeadHREF):           HeadHREF,
	string(HeadLabel):          HeadLabel,
	string(HeadPort):           HeadPort,
	string(HeadTarget):         HeadTarget,
	string(HeadTooltip):        HeadTooltip,
	string(Height):             Height,
	string(HREF):               HREF,
	string(ID):                 ID,
	string(Image):              Image,
	string(ImagePath):          ImagePath,
	string(ImageScale):         ImageScale,
	string(InputScale):         InputScale,
	string(Label):              Label,
	string(LabelURL):           LabelURL,
	string(LabelScheme):        LabelScheme,
	string(LabelAngle):         LabelAngle,
	string(LabelDistance):      LabelDistance,
	string(LabelFloat):         LabelFloat,
	string(LabelFontColor):     LabelFontColor,
	string(LabelFontName):      LabelFontName,
	string(LabelFontSize):      LabelFontSize,
	string(LabelHREF):          LabelHREF,
	string(LabelJust):          LabelJust,
	string(LabelLOC):           LabelLOC,
	string(LabelTarget):        LabelTarget,
	string(LabelTooltip):       LabelTooltip,
	string(Landscape):          Landscape,
	string(Layer):              Layer,
	string(LayerListSep):       LayerListSep,
	string(Layers):             Layers,
	string(LayerSelect):        LayerSelect,
	string(LayerSep):           LayerSep,
	string(Layout):             Layout,
	string(Len):                Len,
	string(Levels):             Levels,
	string(LevelsGap):          LevelsGap,
	string(LHead):              LHead,
	string(LHeight):            LHeight,
	string(LP):                 LP,
	string(LTail):              LTail,
	string(LWidth):             LWidth,
	string(Margin):             Margin,
	string(MaxIter):            MaxIter,
	string(MCLimit):            MCLimit,
	string(MinDist):            MinDist,
	string(MinLen):             MinLen,
	string(Mode):               Mode,
	string(Model):              Model,
	string(Mosek):              Mosek,
	string(NodeSep):            NodeSep,
	string(NoJustify):          NoJustify,
	string(Normalize):          Normalize,
	string(NoTranslate):        NoTranslate,
	string(NSLimit):            NSLimit,
	string(NSLimit1):           NSLimit1,
	string(Ordering):           Ordering,
	string(Orientation):        Orientation,
	string(Outputorder):        Outputorder,
	string(Overlap):            Overlap,
	string(OverlapScaling):     OverlapScaling,
	string(OverlapShrink):      OverlapShrink,
	string(Pack):               Pack,
	string(PackMode):           PackMode,
	string(Pad):                Pad,
	string(Page):               Page,
	string(PageDir):            PageDir,
	string(PenColor):           PenColor,
	string(PenWidth):           PenWidth,
	string(Peripheries):        Peripheries,
	string(Pin):                Pin,
	string(Pos):                Pos,
	string(Quadtree):           Quadtree,
	string(Quantum):            Quantum,
	string(Rank):               Rank,
	string(RankDir):            RankDir,
	string(RankSep):            RankSep,
	string(Ratio):              Ratio,
	string(Rects):              Rects,
	string(Regular):            Regular,
	string(ReMinCross):         ReMinCross,
	string(RepulsiveForce):     RepulsiveForce,
	string(Resolution):         Resolution,
	string(Root):               Root,
	string(Rotate):             Rotate,
	string(Rotation):           Rotation,
	string(SameHead):           SameHead,
	string(SameTail):           SameTail,
	string(SamplePoints):       SamplePoints,
	string(Scale):              Scale,
	string(SearchSize):         SearchSize,
	string(Sep):                Sep,
	string(Shape):              Shape,
	string(ShapeFile):          ShapeFile,
	string(ShowBoxes):          ShowBoxes,
	string(Sides):              Sides,
	string(Size):               Size,
	string(Skew):               Skew,
	string(Smoothing):          Smoothing,
	string(SortV):              SortV,
	string(Splines):            Splines,
	string(Start):              Start,
	string(Style):              Style,
	string(StyleSheet):         StyleSheet,
	string(TailURL):            TailURL,
	string(TailLP):             TailLP,
	string(TailClip):           TailClip,
	string(TailHREF):           TailHREF,
	string(TailLabel):          TailLabel,
	string(TailPort):           TailPort,
	string(TailTarget):         TailTarget,
	string(TailTooltip):        TailTooltip,
	string(Target):             Target,
	string(Tooltip):            Tooltip,
	string(TrueColor):          TrueColor,
	string(Vertices):           Vertices,
	string(ViewPort):           ViewPort,
	string(VoroMargin):         VoroMargin,
	string(Weight):             Weight,
	string(Width):              Width,
	string(XDotVersion):        XDotVersion,
	string(XLabel):             XLabel,
	string(XLP):                XLP,
	string(Z):                  Z,

	string(MinCross): MinCross,
	string(SSize):    SSize,
	string(Outline):  Outline,
	string(F):        F,
}
