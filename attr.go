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
	F        Attr = "f"        // not in the documentation, but found in the transparency.gv.txt example
)

var validAttrs = map[string]Attr{
	string(DAMPING):            DAMPING,
	string(K):                  K,
	string(URL):                URL,
	string(_BACKGROUND):        _BACKGROUND,
	string(AREA):               AREA,
	string(ARROWHEAD):          ARROWHEAD,
	string(ARROWSIZE):          ARROWSIZE,
	string(ARROWTAIL):          ARROWTAIL,
	string(BB):                 BB,
	string(BGCOLOR):            BGCOLOR,
	string(CENTER):             CENTER,
	string(CHARSET):            CHARSET,
	string(CLUSTERRANK):        CLUSTERRANK,
	string(COLOR):              COLOR,
	string(COLORSCHEME):        COLORSCHEME,
	string(COMMENT):            COMMENT,
	string(COMPOUND):           COMPOUND,
	string(CONCENTRATE):        CONCENTRATE,
	string(CONSTRAINT):         CONSTRAINT,
	string(DECORATE):           DECORATE,
	string(DEFAULTDIST):        DEFAULTDIST,
	string(DIM):                DIM,
	string(DIMEN):              DIMEN,
	string(DIR):                DIR,
	string(DIREDGECONSTRAINTS): DIREDGECONSTRAINTS,
	string(DISTORTION):         DISTORTION,
	string(DPI):                DPI,
	string(EDGEURL):            EDGEURL,
	string(EDGEHREF):           EDGEHREF,
	string(EDGETARGET):         EDGETARGET,
	string(EDGETOOLTIP):        EDGETOOLTIP,
	string(EPSILON):            EPSILON,
	string(ESEP):               ESEP,
	string(FILLCOLOR):          FILLCOLOR,
	string(FIXEDSIZE):          FIXEDSIZE,
	string(FONTCOLOR):          FONTCOLOR,
	string(FONTNAME):           FONTNAME,
	string(FONTNAMES):          FONTNAMES,
	string(FONTPATH):           FONTPATH,
	string(FONTSIZE):           FONTSIZE,
	string(FORCELABELS):        FORCELABELS,
	string(GRADIENTANGLE):      GRADIENTANGLE,
	string(GROUP):              GROUP,
	string(HEADURL):            HEADURL,
	string(HEAD_LP):            HEAD_LP,
	string(HEADCLIP):           HEADCLIP,
	string(HEADHREF):           HEADHREF,
	string(HEADLABEL):          HEADLABEL,
	string(HEADPORT):           HEADPORT,
	string(HEADTARGET):         HEADTARGET,
	string(HEADTOOLTIP):        HEADTOOLTIP,
	string(HEIGHT):             HEIGHT,
	string(HREF):               HREF,
	string(ID):                 ID,
	string(IMAGE):              IMAGE,
	string(IMAGEPATH):          IMAGEPATH,
	string(IMAGESCALE):         IMAGESCALE,
	string(INPUTSCALE):         INPUTSCALE,
	string(LABEL):              LABEL,
	string(LABELURL):           LABELURL,
	string(LABEL_SCHEME):       LABEL_SCHEME,
	string(LABELANGLE):         LABELANGLE,
	string(LABELDISTANCE):      LABELDISTANCE,
	string(LABELFLOAT):         LABELFLOAT,
	string(LABELFONTCOLOR):     LABELFONTCOLOR,
	string(LABELFONTNAME):      LABELFONTNAME,
	string(LABELFONTSIZE):      LABELFONTSIZE,
	string(LABELHREF):          LABELHREF,
	string(LABELJUST):          LABELJUST,
	string(LABELLOC):           LABELLOC,
	string(LABELTARGET):        LABELTARGET,
	string(LABELTOOLTIP):       LABELTOOLTIP,
	string(LANDSCAPE):          LANDSCAPE,
	string(LAYER):              LAYER,
	string(LAYERLISTSEP):       LAYERLISTSEP,
	string(LAYERS):             LAYERS,
	string(LAYERSELECT):        LAYERSELECT,
	string(LAYERSEP):           LAYERSEP,
	string(LAYOUT):             LAYOUT,
	string(LEN):                LEN,
	string(LEVELS):             LEVELS,
	string(LEVELSGAP):          LEVELSGAP,
	string(LHEAD):              LHEAD,
	string(LHEIGHT):            LHEIGHT,
	string(LP):                 LP,
	string(LTAIL):              LTAIL,
	string(LWIDTH):             LWIDTH,
	string(MARGIN):             MARGIN,
	string(MAXITER):            MAXITER,
	string(MCLIMIT):            MCLIMIT,
	string(MINDIST):            MINDIST,
	string(MINLEN):             MINLEN,
	string(MODE):               MODE,
	string(MODEL):              MODEL,
	string(MOSEK):              MOSEK,
	string(NODESEP):            NODESEP,
	string(NOJUSTIFY):          NOJUSTIFY,
	string(NORMALIZE):          NORMALIZE,
	string(NOTRANSLATE):        NOTRANSLATE,
	string(NSLIMIT):            NSLIMIT,
	string(NSLIMIT1):           NSLIMIT1,
	string(ORDERING):           ORDERING,
	string(ORIENTATION):        ORIENTATION,
	string(OUTPUTORDER):        OUTPUTORDER,
	string(OVERLAP):            OVERLAP,
	string(OVERLAP_SCALING):    OVERLAP_SCALING,
	string(OVERLAP_SHRINK):     OVERLAP_SHRINK,
	string(PACK):               PACK,
	string(PACKMODE):           PACKMODE,
	string(PAD):                PAD,
	string(PAGE):               PAGE,
	string(PAGEDIR):            PAGEDIR,
	string(PENCOLOR):           PENCOLOR,
	string(PENWIDTH):           PENWIDTH,
	string(PERIPHERIES):        PERIPHERIES,
	string(PIN):                PIN,
	string(POS):                POS,
	string(QUADTREE):           QUADTREE,
	string(QUANTUM):            QUANTUM,
	string(RANK):               RANK,
	string(RANKDIR):            RANKDIR,
	string(RANKSEP):            RANKSEP,
	string(RATIO):              RATIO,
	string(RECTS):              RECTS,
	string(REGULAR):            REGULAR,
	string(REMINCROSS):         REMINCROSS,
	string(REPULSIVEFORCE):     REPULSIVEFORCE,
	string(RESOLUTION):         RESOLUTION,
	string(ROOT):               ROOT,
	string(ROTATE):             ROTATE,
	string(ROTATION):           ROTATION,
	string(SAMEHEAD):           SAMEHEAD,
	string(SAMETAIL):           SAMETAIL,
	string(SAMPLEPOINTS):       SAMPLEPOINTS,
	string(SCALE):              SCALE,
	string(SEARCHSIZE):         SEARCHSIZE,
	string(SEP):                SEP,
	string(SHAPE):              SHAPE,
	string(SHAPEFILE):          SHAPEFILE,
	string(SHOWBOXES):          SHOWBOXES,
	string(SIDES):              SIDES,
	string(SIZE):               SIZE,
	string(SKEW):               SKEW,
	string(SMOOTHING):          SMOOTHING,
	string(SORTV):              SORTV,
	string(SPLINES):            SPLINES,
	string(START):              START,
	string(STYLE):              STYLE,
	string(STYLESHEET):         STYLESHEET,
	string(TAILURL):            TAILURL,
	string(TAIL_LP):            TAIL_LP,
	string(TAILCLIP):           TAILCLIP,
	string(TAILHREF):           TAILHREF,
	string(TAILLABEL):          TAILLABEL,
	string(TAILPORT):           TAILPORT,
	string(TAILTARGET):         TAILTARGET,
	string(TAILTOOLTIP):        TAILTOOLTIP,
	string(TARGET):             TARGET,
	string(TOOLTIP):            TOOLTIP,
	string(TRUECOLOR):          TRUECOLOR,
	string(VERTICES):           VERTICES,
	string(VIEWPORT):           VIEWPORT,
	string(VORO_MARGIN):        VORO_MARGIN,
	string(WEIGHT):             WEIGHT,
	string(WIDTH):              WIDTH,
	string(XDOTVERSION):        XDOTVERSION,
	string(XLABEL):             XLABEL,
	string(XLP):                XLP,
	string(Z):                  Z,

	string(MINCROSS): MINCROSS,
	string(SSIZE):    SSIZE,
	string(OUTLINE):  OUTLINE,
	string(F):        F,
}
