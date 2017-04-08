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

type Attr string

const (
	DAMPING            Attr = "Damping"
	K                  Attr = "K"
	URL                Attr = "URL"
	_BACKGROUND        Attr = "_background"
	AREA               Attr = "area"
	ARROWHEAD          Attr = "arrowhead"
	ARROWSIZE          Attr = "arrowsize"
	ARROWTAIL          Attr = "arrowtail"
	BB                 Attr = "bb"
	BGCOLOR            Attr = "bgcolor"
	CENTER             Attr = "center"
	CHARSET            Attr = "charset"
	CLUSTERRANK        Attr = "clusterrank"
	COLOR              Attr = "color"
	COLORSCHEME        Attr = "colorscheme"
	COMMENT            Attr = "comment"
	COMPOUND           Attr = "compound"
	CONCENTRATE        Attr = "concentrate"
	CONSTRAINT         Attr = "constraint"
	DECORATE           Attr = "decorate"
	DEFAULTDIST        Attr = "defaultdist"
	DIM                Attr = "dim"
	DIMEN              Attr = "dimen"
	DIR                Attr = "dir"
	DIREDGECONSTRAINTS Attr = "diredgeconstraints"
	DISTORTION         Attr = "distortion"
	DPI                Attr = "dpi"
	EDGEURL            Attr = "edgeURL"
	EDGEHREF           Attr = "edgehref"
	EDGETARGET         Attr = "edgetarget"
	EDGETOOLTIP        Attr = "edgetooltip"
	EPSILON            Attr = "epsilon"
	ESEP               Attr = "esep"
	FILLCOLOR          Attr = "fillcolor"
	FIXEDSIZE          Attr = "fixedsize"
	FONTCOLOR          Attr = "fontcolor"
	FONTNAME           Attr = "fontname"
	FONTNAMES          Attr = "fontnames"
	FONTPATH           Attr = "fontpath"
	FONTSIZE           Attr = "fontsize"
	FORCELABELS        Attr = "forcelabels"
	GRADIENTANGLE      Attr = "gradientangle"
	GROUP              Attr = "group"
	HEADURL            Attr = "headURL"
	HEAD_LP            Attr = "head_lp"
	HEADCLIP           Attr = "headclip"
	HEADHREF           Attr = "headhref"
	HEADLABEL          Attr = "headlabel"
	HEADPORT           Attr = "headport"
	HEADTARGET         Attr = "headtarget"
	HEADTOOLTIP        Attr = "headtooltip"
	HEIGHT             Attr = "height"
	HREF               Attr = "href"
	ID                 Attr = "id"
	IMAGE              Attr = "image"
	IMAGEPATH          Attr = "imagepath"
	IMAGESCALE         Attr = "imagescale"
	INPUTSCALE         Attr = "inputscale"
	LABEL              Attr = "label"
	LABELURL           Attr = "labelURL"
	LABEL_SCHEME       Attr = "label_scheme"
	LABELANGLE         Attr = "labelangle"
	LABELDISTANCE      Attr = "labeldistance"
	LABELFLOAT         Attr = "labelfloat"
	LABELFONTCOLOR     Attr = "labelfontcolor"
	LABELFONTNAME      Attr = "labelfontname"
	LABELFONTSIZE      Attr = "labelfontsize"
	LABELHREF          Attr = "labelhref"
	LABELJUST          Attr = "labeljust"
	LABELLOC           Attr = "labelloc"
	LABELTARGET        Attr = "labeltarget"
	LABELTOOLTIP       Attr = "labeltooltip"
	LANDSCAPE          Attr = "landscape"
	LAYER              Attr = "layer"
	LAYERLISTSEP       Attr = "layerlistsep"
	LAYERS             Attr = "layers"
	LAYERSELECT        Attr = "layerselect"
	LAYERSEP           Attr = "layersep"
	LAYOUT             Attr = "layout"
	LEN                Attr = "len"
	LEVELS             Attr = "levels"
	LEVELSGAP          Attr = "levelsgap"
	LHEAD              Attr = "lhead"
	LHEIGHT            Attr = "lheight"
	LP                 Attr = "lp"
	LTAIL              Attr = "ltail"
	LWIDTH             Attr = "lwidth"
	MARGIN             Attr = "margin"
	MAXITER            Attr = "maxiter"
	MCLIMIT            Attr = "mclimit"
	MINDIST            Attr = "mindist"
	MINLEN             Attr = "minlen"
	MODE               Attr = "mode"
	MODEL              Attr = "model"
	MOSEK              Attr = "mosek"
	NODESEP            Attr = "nodesep"
	NOJUSTIFY          Attr = "nojustify"
	NORMALIZE          Attr = "normalize"
	NOTRANSLATE        Attr = "notranslate"
	NSLIMIT            Attr = "nslimit"
	NSLIMIT1           Attr = "nslimit1"
	ORDERING           Attr = "ordering"
	ORIENTATION        Attr = "orientation"
	OUTPUTORDER        Attr = "outputorder"
	OVERLAP            Attr = "overlap"
	OVERLAP_SCALING    Attr = "overlap_scaling"
	OVERLAP_SHRINK     Attr = "overlap_shrink"
	PACK               Attr = "pack"
	PACKMODE           Attr = "packmode"
	PAD                Attr = "pad"
	PAGE               Attr = "page"
	PAGEDIR            Attr = "pagedir"
	PENCOLOR           Attr = "pencolor"
	PENWIDTH           Attr = "penwidth"
	PERIPHERIES        Attr = "peripheries"
	PIN                Attr = "pin"
	POS                Attr = "pos"
	QUADTREE           Attr = "quadtree"
	QUANTUM            Attr = "quantum"
	RANK               Attr = "rank"
	RANKDIR            Attr = "rankdir"
	RANKSEP            Attr = "ranksep"
	RATIO              Attr = "ratio"
	RECTS              Attr = "rects"
	REGULAR            Attr = "regular"
	REMINCROSS         Attr = "remincross"
	REPULSIVEFORCE     Attr = "repulsiveforce"
	RESOLUTION         Attr = "resolution"
	ROOT               Attr = "root"
	ROTATE             Attr = "rotate"
	ROTATION           Attr = "rotation"
	SAMEHEAD           Attr = "samehead"
	SAMETAIL           Attr = "sametail"
	SAMPLEPOINTS       Attr = "samplepoints"
	SCALE              Attr = "scale"
	SEARCHSIZE         Attr = "searchsize"
	SEP                Attr = "sep"
	SHAPE              Attr = "shape"
	SHAPEFILE          Attr = "shapefile"
	SHOWBOXES          Attr = "showboxes"
	SIDES              Attr = "sides"
	SIZE               Attr = "size"
	SKEW               Attr = "skew"
	SMOOTHING          Attr = "smoothing"
	SORTV              Attr = "sortv"
	SPLINES            Attr = "splines"
	START              Attr = "start"
	STYLE              Attr = "style"
	STYLESHEET         Attr = "stylesheet"
	TAILURL            Attr = "tailURL"
	TAIL_LP            Attr = "tail_lp"
	TAILCLIP           Attr = "tailclip"
	TAILHREF           Attr = "tailhref"
	TAILLABEL          Attr = "taillabel"
	TAILPORT           Attr = "tailport"
	TAILTARGET         Attr = "tailtarget"
	TAILTOOLTIP        Attr = "tailtooltip"
	TARGET             Attr = "target"
	TOOLTIP            Attr = "tooltip"
	TRUECOLOR          Attr = "truecolor"
	VERTICES           Attr = "vertices"
	VIEWPORT           Attr = "viewport"
	VORO_MARGIN        Attr = "voro_margin"
	WEIGHT             Attr = "weight"
	WIDTH              Attr = "width"
	XDOTVERSION        Attr = "xdotversion"
	XLABEL             Attr = "xlabel"
	XLP                Attr = "xlp"
	Z                  Attr = "z"

	MINCROSS Attr = "mincross" // not in the documentation, but found in the Ped_Lion_Share (lion_share.gv.txt) example
	SSIZE    Attr = "ssize"    // not in the documentation, but found in the siblings.gv.txt example
	OUTLINE  Attr = "outline"  // not in the documentation, but found in the siblings.gv.txt example
)

var validAttrs = map[string]struct{}{
	string(DAMPING):            {},
	string(K):                  {},
	string(URL):                {},
	string(_BACKGROUND):        {},
	string(AREA):               {},
	string(ARROWHEAD):          {},
	string(ARROWSIZE):          {},
	string(ARROWTAIL):          {},
	string(BB):                 {},
	string(BGCOLOR):            {},
	string(CENTER):             {},
	string(CHARSET):            {},
	string(CLUSTERRANK):        {},
	string(COLOR):              {},
	string(COLORSCHEME):        {},
	string(COMMENT):            {},
	string(COMPOUND):           {},
	string(CONCENTRATE):        {},
	string(CONSTRAINT):         {},
	string(DECORATE):           {},
	string(DEFAULTDIST):        {},
	string(DIM):                {},
	string(DIMEN):              {},
	string(DIR):                {},
	string(DIREDGECONSTRAINTS): {},
	string(DISTORTION):         {},
	string(DPI):                {},
	string(EDGEURL):            {},
	string(EDGEHREF):           {},
	string(EDGETARGET):         {},
	string(EDGETOOLTIP):        {},
	string(EPSILON):            {},
	string(ESEP):               {},
	string(FILLCOLOR):          {},
	string(FIXEDSIZE):          {},
	string(FONTCOLOR):          {},
	string(FONTNAME):           {},
	string(FONTNAMES):          {},
	string(FONTPATH):           {},
	string(FONTSIZE):           {},
	string(FORCELABELS):        {},
	string(GRADIENTANGLE):      {},
	string(GROUP):              {},
	string(HEADURL):            {},
	string(HEAD_LP):            {},
	string(HEADCLIP):           {},
	string(HEADHREF):           {},
	string(HEADLABEL):          {},
	string(HEADPORT):           {},
	string(HEADTARGET):         {},
	string(HEADTOOLTIP):        {},
	string(HEIGHT):             {},
	string(HREF):               {},
	string(ID):                 {},
	string(IMAGE):              {},
	string(IMAGEPATH):          {},
	string(IMAGESCALE):         {},
	string(INPUTSCALE):         {},
	string(LABEL):              {},
	string(LABELURL):           {},
	string(LABEL_SCHEME):       {},
	string(LABELANGLE):         {},
	string(LABELDISTANCE):      {},
	string(LABELFLOAT):         {},
	string(LABELFONTCOLOR):     {},
	string(LABELFONTNAME):      {},
	string(LABELFONTSIZE):      {},
	string(LABELHREF):          {},
	string(LABELJUST):          {},
	string(LABELLOC):           {},
	string(LABELTARGET):        {},
	string(LABELTOOLTIP):       {},
	string(LANDSCAPE):          {},
	string(LAYER):              {},
	string(LAYERLISTSEP):       {},
	string(LAYERS):             {},
	string(LAYERSELECT):        {},
	string(LAYERSEP):           {},
	string(LAYOUT):             {},
	string(LEN):                {},
	string(LEVELS):             {},
	string(LEVELSGAP):          {},
	string(LHEAD):              {},
	string(LHEIGHT):            {},
	string(LP):                 {},
	string(LTAIL):              {},
	string(LWIDTH):             {},
	string(MARGIN):             {},
	string(MAXITER):            {},
	string(MCLIMIT):            {},
	string(MINDIST):            {},
	string(MINLEN):             {},
	string(MODE):               {},
	string(MODEL):              {},
	string(MOSEK):              {},
	string(NODESEP):            {},
	string(NOJUSTIFY):          {},
	string(NORMALIZE):          {},
	string(NOTRANSLATE):        {},
	string(NSLIMIT):            {},
	string(NSLIMIT1):           {},
	string(ORDERING):           {},
	string(ORIENTATION):        {},
	string(OUTPUTORDER):        {},
	string(OVERLAP):            {},
	string(OVERLAP_SCALING):    {},
	string(OVERLAP_SHRINK):     {},
	string(PACK):               {},
	string(PACKMODE):           {},
	string(PAD):                {},
	string(PAGE):               {},
	string(PAGEDIR):            {},
	string(PENCOLOR):           {},
	string(PENWIDTH):           {},
	string(PERIPHERIES):        {},
	string(PIN):                {},
	string(POS):                {},
	string(QUADTREE):           {},
	string(QUANTUM):            {},
	string(RANK):               {},
	string(RANKDIR):            {},
	string(RANKSEP):            {},
	string(RATIO):              {},
	string(RECTS):              {},
	string(REGULAR):            {},
	string(REMINCROSS):         {},
	string(REPULSIVEFORCE):     {},
	string(RESOLUTION):         {},
	string(ROOT):               {},
	string(ROTATE):             {},
	string(ROTATION):           {},
	string(SAMEHEAD):           {},
	string(SAMETAIL):           {},
	string(SAMPLEPOINTS):       {},
	string(SCALE):              {},
	string(SEARCHSIZE):         {},
	string(SEP):                {},
	string(SHAPE):              {},
	string(SHAPEFILE):          {},
	string(SHOWBOXES):          {},
	string(SIDES):              {},
	string(SIZE):               {},
	string(SKEW):               {},
	string(SMOOTHING):          {},
	string(SORTV):              {},
	string(SPLINES):            {},
	string(START):              {},
	string(STYLE):              {},
	string(STYLESHEET):         {},
	string(TAILURL):            {},
	string(TAIL_LP):            {},
	string(TAILCLIP):           {},
	string(TAILHREF):           {},
	string(TAILLABEL):          {},
	string(TAILPORT):           {},
	string(TAILTARGET):         {},
	string(TAILTOOLTIP):        {},
	string(TARGET):             {},
	string(TOOLTIP):            {},
	string(TRUECOLOR):          {},
	string(VERTICES):           {},
	string(VIEWPORT):           {},
	string(VORO_MARGIN):        {},
	string(WEIGHT):             {},
	string(WIDTH):              {},
	string(XDOTVERSION):        {},
	string(XLABEL):             {},
	string(XLP):                {},
	string(Z):                  {},

	string(MINCROSS): {},
	string(SSIZE):    {},
	string(OUTLINE):  {},
}
