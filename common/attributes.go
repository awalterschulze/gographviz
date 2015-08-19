//Copyright 2013 Vastech SA (PTY) LTD
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package common

type Attribute int

const (
	DAMPING Attribute = iota
	K
	URL
	_BACKGROUND
	AREA
	ARROWHEAD
	ARROWSIZE
	ARROWTAIL
	BB
	BGCOLOR
	CENTER
	CHARSET
	CLUSTERRANK
	COLOR
	COLORSCHEME
	COMMENT
	COMPOUND
	CONCENTRATE
	CONSTRAINT
	DECORATE
	DEFAULTDIST
	DIM
	DIMEN
	DIR
	DIREDGECONSTRAINTS
	DISTORTION
	DPI
	EDGEURL
	EDGEHREF
	EDGETARGET
	EDGETOOLTIP
	EPSILON
	ESEP
	FILLCOLOR
	FIXEDSIZE
	FONTCOLOR
	FONTNAME
	FONTNAMES
	FONTPATH
	FONTSIZE
	FORCELABELS
	GRADIENTANGLE
	GROUP
	HEADURL
	HEAD_LP
	HEADCLIP
	HEADHREF
	HEADLABEL
	HEADPORT
	HEADTARGET
	HEADTOOLTIP
	HEIGHT
	HREF
	ID
	IMAGE
	IMAGEPATH
	IMAGESCALE
	INPUTSCALE
	LABEL
	LABELURL
	LABEL_SCHEME
	LABELANGLE
	LABELDISTANCE
	LABELFLOAT
	LABELFONTCOLOR
	LABELFONTNAME
	LABELFONTSIZE
	LABELHREF
	LABELJUST
	LABELLOC
	LABELTARGET
	LABELTOOLTIP
	LANDSCAPE
	LAYER
	LAYERLISTSEP
	LAYERS
	LAYERSELECT
	LAYERSEP
	LAYOUT
	LEN
	LEVELS
	LEVELSGAP
	LHEAD
	LHEIGHT
	LP
	LTAIL
	LWIDTH
	MARGIN
	MAXITER
	MCLIMIT
	MINDIST
	MINLEN
	MODE
	MODEL
	MOSEK
	NODESEP
	NOJUSTIFY
	NORMALIZE
	NOTRANSLATE
	NSLIMIT
	NSLIMIT1
	ORDERING
	ORIENTATION
	OUTPUTORDER
	OVERLAP
	OVERLAP_SCALING
	OVERLAP_SHRINK
	PACK
	PACKMODE
	PAD
	PAGE
	PAGEDIR
	PENCOLOR
	PENWIDTH
	PERIPHERIES
	PIN
	POS
	QUADTREE
	QUANTUM
	RANK
	RANKDIR
	RANKSEP
	RATIO
	RECTS
	REGULAR
	REMINCROSS
	REPULSIVEFORCE
	RESOLUTION
	ROOT
	ROTATE
	ROTATION
	SAMEHEAD
	SAMETAIL
	SAMPLEPOINTS
	SCALE
	SEARCHSIZE
	SEP
	SHAPE
	SHAPEFILE
	SHOWBOXES
	SIDES
	SIZE
	SKEW
	SMOOTHING
	SORTV
	SPLINES
	START
	STYLE
	STYLESHEET
	TAILURL
	TAIL_LP
	TAILCLIP
	TAILHREF
	TAILLABEL
	TAILPORT
	TAILTARGET
	TAILTOOLTIP
	TARGET
	TOOLTIP
	TRUECOLOR
	VERTICES
	VIEWPORT
	VORO_MARGIN
	WEIGHT
	WIDTH
	XDOTVERSION
	XLABEL
	XLP
	Z

	MAXATTRIBUTE
)

var takesString [MAXATTRIBUTE]bool = [MAXATTRIBUTE]bool{
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	true,
	true,
	false,
	true,
	false,
	false,
	false,
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	true,
	true,
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	true,
	false,
	true,
	true,
	true,
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	false,
	true,
	false,
	true,
	true,
	true,
	false,
	true,
	false,
	true,
	true,
	true,
	true,
	true,
	true,
	false,
	false,
	false,
	true,
	false,
	false,
	true,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	true,
	false,
	true,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
	true,
	true,
	true,
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	true,
	false,
	true,
	false,
	false,
	true,
	false,
	false,
	false,
	false,
	true,
	false,
	false,
}

var stringLookup map[string]Attribute = map[string]Attribute{
	"Damping":            DAMPING,
	"K":                  K,
	"URL":                URL,
	"_background":        _BACKGROUND,
	"area":               AREA,
	"arrowhead":          ARROWHEAD,
	"arrowsize":          ARROWSIZE,
	"arrowtail":          ARROWTAIL,
	"bb":                 BB,
	"bgcolor":            BGCOLOR,
	"center":             CENTER,
	"charset":            CHARSET,
	"clusterrank":        CLUSTERRANK,
	"color":              COLOR,
	"colorscheme":        COLORSCHEME,
	"comment":            COMMENT,
	"compound":           COMPOUND,
	"concentrate":        CONCENTRATE,
	"constraint":         CONSTRAINT,
	"decorate":           DECORATE,
	"defaultdist":        DEFAULTDIST,
	"dim":                DIM,
	"dimen":              DIMEN,
	"dir":                DIR,
	"diredgeconstraints": DIREDGECONSTRAINTS,
	"distortion":         DISTORTION,
	"dpi":                DPI,
	"edgeURL":            EDGEURL,
	"edgehref":           EDGEHREF,
	"edgetarget":         EDGETARGET,
	"edgetooltip":        EDGETOOLTIP,
	"epsilon":            EPSILON,
	"esep":               ESEP,
	"fillcolor":          FILLCOLOR,
	"fixedsize":          FIXEDSIZE,
	"fontcolor":          FONTCOLOR,
	"fontname":           FONTNAME,
	"fontnames":          FONTNAMES,
	"fontpath":           FONTPATH,
	"fontsize":           FONTSIZE,
	"forcelabels":        FORCELABELS,
	"gradientangle":      GRADIENTANGLE,
	"group":              GROUP,
	"headURL":            HEADURL,
	"head_lp":            HEAD_LP,
	"headclip":           HEADCLIP,
	"headhref":           HEADHREF,
	"headlabel":          HEADLABEL,
	"headport":           HEADPORT,
	"headtarget":         HEADTARGET,
	"headtooltip":        HEADTOOLTIP,
	"height":             HEIGHT,
	"href":               HREF,
	"id":                 ID,
	"image":              IMAGE,
	"imagepath":          IMAGEPATH,
	"imagescale":         IMAGESCALE,
	"inputscale":         INPUTSCALE,
	"label":              LABEL,
	"labelURL":           LABELURL,
	"label_scheme":       LABEL_SCHEME,
	"labelangle":         LABELANGLE,
	"labeldistance":      LABELDISTANCE,
	"labelfloat":         LABELFLOAT,
	"labelfontcolor":     LABELFONTCOLOR,
	"labelfontname":      LABELFONTNAME,
	"labelfontsize":      LABELFONTSIZE,
	"labelhref":          LABELHREF,
	"labeljust":          LABELJUST,
	"labelloc":           LABELLOC,
	"labeltarget":        LABELTARGET,
	"labeltooltip":       LABELTOOLTIP,
	"landscape":          LANDSCAPE,
	"layer":              LAYER,
	"layerlistsep":       LAYERLISTSEP,
	"layers":             LAYERS,
	"layerselect":        LAYERSELECT,
	"layersep":           LAYERSEP,
	"layout":             LAYOUT,
	"len":                LEN,
	"levels":             LEVELS,
	"levelsgap":          LEVELSGAP,
	"lhead":              LHEAD,
	"lheight":            LHEIGHT,
	"lp":                 LP,
	"ltail":              LTAIL,
	"lwidth":             LWIDTH,
	"margin":             MARGIN,
	"maxiter":            MAXITER,
	"mclimit":            MCLIMIT,
	"mindist":            MINDIST,
	"minlen":             MINLEN,
	"mode":               MODE,
	"model":              MODEL,
	"mosek":              MOSEK,
	"nodesep":            NODESEP,
	"nojustify":          NOJUSTIFY,
	"normalize":          NORMALIZE,
	"notranslate":        NOTRANSLATE,
	"nslimit":            NSLIMIT,
	"nslimit1":           NSLIMIT1,
	"ordering":           ORDERING,
	"orientation":        ORIENTATION,
	"outputorder":        OUTPUTORDER,
	"overlap":            OVERLAP,
	"overlap_scaling":    OVERLAP_SCALING,
	"overlap_shrink":     OVERLAP_SHRINK,
	"pack":               PACK,
	"packmode":           PACKMODE,
	"pad":                PAD,
	"page":               PAGE,
	"pagedir":            PAGEDIR,
	"pencolor":           PENCOLOR,
	"penwidth":           PENWIDTH,
	"peripheries":        PERIPHERIES,
	"pin":                PIN,
	"pos":                POS,
	"quadtree":           QUADTREE,
	"quantum":            QUANTUM,
	"rank":               RANK,
	"rankdir":            RANKDIR,
	"ranksep":            RANKSEP,
	"ratio":              RATIO,
	"rects":              RECTS,
	"regular":            REGULAR,
	"remincross":         REMINCROSS,
	"repulsiveforce":     REPULSIVEFORCE,
	"resolution":         RESOLUTION,
	"root":               ROOT,
	"rotate":             ROTATE,
	"rotation":           ROTATION,
	"samehead":           SAMEHEAD,
	"sametail":           SAMETAIL,
	"samplepoints":       SAMPLEPOINTS,
	"scale":              SCALE,
	"searchsize":         SEARCHSIZE,
	"sep":                SEP,
	"shape":              SHAPE,
	"shapefile":          SHAPEFILE,
	"showboxes":          SHOWBOXES,
	"sides":              SIDES,
	"size":               SIZE,
	"skew":               SKEW,
	"smoothing":          SMOOTHING,
	"sortv":              SORTV,
	"splines":            SPLINES,
	"start":              START,
	"style":              STYLE,
	"stylesheet":         STYLESHEET,
	"tailURL":            TAILURL,
	"tail_lp":            TAIL_LP,
	"tailclip":           TAILCLIP,
	"tailhref":           TAILHREF,
	"taillabel":          TAILLABEL,
	"tailport":           TAILPORT,
	"tailtarget":         TAILTARGET,
	"tailtooltip":        TAILTOOLTIP,
	"target":             TARGET,
	"tooltip":            TOOLTIP,
	"truecolor":          TRUECOLOR,
	"vertices":           VERTICES,
	"viewport":           VIEWPORT,
	"voro_margin":        VORO_MARGIN,
	"weight":             WEIGHT,
	"width":              WIDTH,
	"xdotversion":        XDOTVERSION,
	"xlabel":             XLABEL,
	"xlp":                XLP,
	"z":                  Z,
}

func (a Attribute) TakesString() bool {
	return takesString[a]
}

func StringToAttribute(input string) (Attribute, bool) {
	attr, ok := stringLookup[input]
	return attr, ok
}
