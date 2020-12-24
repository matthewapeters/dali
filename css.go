package dali

import (
	"fmt"
	"sort"
)

// Property descriptions originally sourced from the following sites
// Content in the form of comments has been moderately modified from original work

// https://www.w3schools.com/cssref/default.asp
// Copyright 1999-2020 by Refsnes Data. All Rights Reserved.
//

/*
W3C Properties Approved For Everyday Use:

Ab­bre­vi­a­tion	Full name
ED		Editors' Draft (not a W3C Technical Report)
FPWD	First Public Working Draft
WD		Working Draft

CR		Candidate Recommendation				--------+
CRD		Candidate Recommendation Draft                  | // Encouraged everyday use
PR		Proposed Recommendation                         |
REC		Recommendation							--------+

NOTE	Working Group Note

The names are defined in sections 6.2.1 and 6.3 of the W3C process document.
A REC is what is normally referred to as a “standard.”
W3C encourages everyday use starting from CR.

Pub.	Property	Specification	St.
--*	CSS Custom Properties for Cascading Variables Module Level 1	CR
align-content	CSS Flexible Box Layout Module Level 1	CR
align-items	CSS Flexible Box Layout Module Level 1	CR
align-self	CSS Flexible Box Layout Module Level 1	CR
all	CSS Cascading and Inheritance Level 3	PR
azimuth	CSS 2.1	REC
background	CSS 2.1	REC
background	CSS Backgrounds and Borders Module Level 3	CR
background-attachment	CSS 2.1	REC
background-attachment	CSS Backgrounds and Borders Module Level 3	CR
background-blend-mode	Compositing and Blending Level 1	CR
background-clip	CSS Backgrounds and Borders Module Level 3	CR
background-color	CSS 2.1	REC
background-color	CSS Backgrounds and Borders Module Level 3	CR
background-image	CSS 2.1	REC
background-image	CSS Backgrounds and Borders Module Level 3	CR
background-origin	CSS Backgrounds and Borders Module Level 3	CR
background-position	CSS 2.1	REC
background-position	CSS Backgrounds and Borders Module Level 3	CR
background-repeat	CSS 2.1	REC
background-repeat	CSS Backgrounds and Borders Module Level 3	CR
background-size	CSS Backgrounds and Borders Module Level 3	CR
border	CSS 2.1	REC
border	CSS Backgrounds and Borders Module Level 3	CR
border-bottom	CSS 2.1	REC
border-bottom	CSS Backgrounds and Borders Module Level 3	CR
border-bottom-color	CSS 2.1	REC
border-bottom-color	CSS Backgrounds and Borders Module Level 3	CR
border-bottom-left-radius	CSS Backgrounds and Borders Module Level 3	CR
border-bottom-right-radius	CSS Backgrounds and Borders Module Level 3	CR
border-bottom-style	CSS 2.1	REC
border-bottom-style	CSS Backgrounds and Borders Module Level 3	CR
border-bottom-width	CSS 2.1	REC
border-bottom-width	CSS Backgrounds and Borders Module Level 3	CR
border-collapse	CSS 2.1	REC
border-color	CSS 2.1	REC
border-color	CSS Backgrounds and Borders Module Level 3	CR
border-image	CSS Backgrounds and Borders Module Level 3	CR
border-image-outset	CSS Backgrounds and Borders Module Level 3	CR
border-image-repeat	CSS Backgrounds and Borders Module Level 3	CR
border-image-slice	CSS Backgrounds and Borders Module Level 3	CR
border-image-source	CSS Backgrounds and Borders Module Level 3	CR
border-image-width	CSS Backgrounds and Borders Module Level 3	CR
border-left	CSS 2.1	REC
border-left	CSS Backgrounds and Borders Module Level 3	CR
border-left-color	CSS 2.1	REC
border-left-color	CSS Backgrounds and Borders Module Level 3	CR
border-left-style	CSS 2.1	REC
border-left-style	CSS Backgrounds and Borders Module Level 3	CR
border-left-width	CSS 2.1	REC
border-left-width	CSS Backgrounds and Borders Module Level 3	CR
border-radius	CSS Backgrounds and Borders Module Level 3	CR
border-right	CSS 2.1	REC
border-right	CSS Backgrounds and Borders Module Level 3	CR
border-right-color	CSS 2.1	REC
border-right-color	CSS Backgrounds and Borders Module Level 3	CR
border-right-style	CSS 2.1	REC
border-right-style	CSS Backgrounds and Borders Module Level 3	CR
border-right-width	CSS 2.1	REC
border-right-width	CSS Backgrounds and Borders Module Level 3	CR
border-spacing	CSS 2.1	REC
border-style	CSS 2.1	REC
border-style	CSS Backgrounds and Borders Module Level 3	CR
border-top	CSS 2.1	REC
border-top	CSS Backgrounds and Borders Module Level 3	CR
border-top-color	CSS 2.1	REC
border-top-color	CSS Backgrounds and Borders Module Level 3	CR
border-top-left-radius	CSS Backgrounds and Borders Module Level 3	CR
border-top-right-radius	CSS Backgrounds and Borders Module Level 3	CR
border-top-style	CSS 2.1	REC
border-top-style	CSS Backgrounds and Borders Module Level 3	CR
border-top-width	CSS 2.1	REC
border-top-width	CSS Backgrounds and Borders Module Level 3	CR
border-width	CSS 2.1	REC
border-width	CSS Backgrounds and Borders Module Level 3	CR
bottom	CSS 2.1	REC
box-decoration-break	CSS Fragmentation Module Level 3	CR
box-shadow	CSS Backgrounds and Borders Module Level 3	CR
box-sizing	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
break-after	CSS Fragmentation Module Level 3	CR
break-before	CSS Fragmentation Module Level 3	CR
break-inside	CSS Fragmentation Module Level 3	CR
caption-side	CSS 2.1	REC
caret-color	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
chains	CSS Template Layout Module	NOTE
clear	CSS 2.1	REC
clip	CSS 2.1	REC
clip	CSS Masking Module Level 1	CR
clip-path	CSS Masking Module Level 1	CR
clip-rule	CSS Masking Module Level 1	CR
color	CSS 2.1	REC
color	CSS Color Module Level 3	REC
contain	CSS Containment Module Level 1	REC
content	CSS 2.1	REC
counter-increment	CSS 2.1	REC
counter-reset	CSS 2.1	REC
cue	CSS 2.1	REC
cue	CSS Speech Module	CR
cue-after	CSS 2.1	REC
cue-after	CSS Speech Module	CR
cue-before	CSS 2.1	REC
cue-before	CSS Speech Module	CR
cursor	CSS 2.1	REC
cursor	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
direction	CSS 2.1	REC
direction	CSS Writing Modes Level 3	REC
direction	CSS Writing Modes Level 4	CR
display	CSS 2.1	REC
display	CSS Display Module Level 3	CR
elevation	CSS 2.1	REC
empty-cells	CSS 2.1	REC
flex	CSS Flexible Box Layout Module Level 1	CR
flex-basis	CSS Flexible Box Layout Module Level 1	CR
flex-direction	CSS Flexible Box Layout Module Level 1	CR
flex-flow	CSS Flexible Box Layout Module Level 1	CR
flex-grow	CSS Flexible Box Layout Module Level 1	CR
flex-shrink	CSS Flexible Box Layout Module Level 1	CR
flex-wrap	CSS Flexible Box Layout Module Level 1	CR
float	CSS 2.1	REC
flow	CSS Template Layout Module	NOTE
font	CSS 2.1	REC
font	CSS Fonts Module Level 3	REC
font-family	CSS 2.1	REC
font-family	CSS Fonts Module Level 3	REC
font-feature-settings	CSS Fonts Module Level 3	REC
font-kerning	CSS Fonts Module Level 3	REC
font-size	CSS 2.1	REC
font-size	CSS Fonts Module Level 3	REC
font-size-adjust	CSS Fonts Module Level 3	REC
font-stretch	CSS Fonts Module Level 3	REC
font-style	CSS 2.1	REC
font-style	CSS Fonts Module Level 3	REC
font-synthesis	CSS Fonts Module Level 3	REC
font-variant	CSS 2.1	REC
font-variant	CSS Fonts Module Level 3	REC
font-variant-caps	CSS Fonts Module Level 3	REC
font-variant-east-asian	CSS Fonts Module Level 3	REC
font-variant-ligatures	CSS Fonts Module Level 3	REC
font-variant-numeric	CSS Fonts Module Level 3	REC
font-variant-position	CSS Fonts Module Level 3	REC
font-weight	CSS 2.1	REC
font-weight	CSS Fonts Module Level 3	REC
glyph-orientation-vertical	CSS Writing Modes Level 3	REC
glyph-orientation-vertical	CSS Writing Modes Level 4	CR
grid	CSS Grid Layout Module Level 1	CR
grid	CSS Grid Layout Module Level 2	CR
grid	CSS Template Layout Module	NOTE
grid-area	CSS Grid Layout Module Level 1	CR
grid-area	CSS Grid Layout Module Level 2	CR
grid-auto-columns	CSS Grid Layout Module Level 1	CR
grid-auto-columns	CSS Grid Layout Module Level 2	CR
grid-auto-flow	CSS Grid Layout Module Level 1	CR
grid-auto-flow	CSS Grid Layout Module Level 2	CR
grid-auto-rows	CSS Grid Layout Module Level 1	CR
grid-auto-rows	CSS Grid Layout Module Level 2	CR
grid-column	CSS Grid Layout Module Level 1	CR
grid-column	CSS Grid Layout Module Level 2	CR
grid-column-end	CSS Grid Layout Module Level 1	CR
grid-column-end	CSS Grid Layout Module Level 2	CR
grid-column-start	CSS Grid Layout Module Level 1	CR
grid-column-start	CSS Grid Layout Module Level 2	CR
grid-row	CSS Grid Layout Module Level 1	CR
grid-row	CSS Grid Layout Module Level 2	CR
grid-row-end	CSS Grid Layout Module Level 1	CR
grid-row-end	CSS Grid Layout Module Level 2	CR
grid-row-start	CSS Grid Layout Module Level 1	CR
grid-row-start	CSS Grid Layout Module Level 2	CR
grid-template	CSS Grid Layout Module Level 1	CR
grid-template	CSS Grid Layout Module Level 2	CR
grid-template	CSS Template Layout Module	NOTE
grid-template-areas	CSS Grid Layout Module Level 1	CR
grid-template-areas	CSS Grid Layout Module Level 2	CR
grid-template-areas	CSS Template Layout Module	NOTE
grid-template-columns	CSS Grid Layout Module Level 1	CR
grid-template-columns	CSS Grid Layout Module Level 2	CR
grid-template-columns	CSS Template Layout Module	NOTE
grid-template-rows	CSS Grid Layout Module Level 1	CR
grid-template-rows	CSS Grid Layout Module Level 2	CR
grid-template-rows	CSS Template Layout Module	NOTE
hanging-punctuation	CSS Text Module Level 3	CR
height	CSS 2.1	REC
hyphens	CSS Text Module Level 3	CR
image-orientation	CSS Images Module Level 3	CR
image-rendering	CSS Images Module Level 3	CR
isolation	Compositing and Blending Level 1	CR
justify-content	CSS Flexible Box Layout Module Level 1	CR
left	CSS 2.1	REC
letter-spacing	CSS 2.1	REC
letter-spacing	CSS Text Module Level 3	CR
line-break	CSS Text Module Level 3	CR
line-height	CSS 2.1	REC
list-style	CSS 2.1	REC
list-style-image	CSS 2.1	REC
list-style-position	CSS 2.1	REC
list-style-type	CSS 2.1	REC
margin	CSS 2.1	REC
margin	CSS Box Model Module Level 3	CR
margin-bottom	CSS 2.1	REC
margin-bottom	CSS Box Model Module Level 3	CR
margin-left	CSS 2.1	REC
margin-left	CSS Box Model Module Level 3	CR
margin-right	CSS 2.1	REC
margin-right	CSS Box Model Module Level 3	CR
margin-top	CSS 2.1	REC
margin-top	CSS Box Model Module Level 3	CR
mask	CSS Masking Module Level 1	CR
mask-border	CSS Masking Module Level 1	CR
mask-border-mode	CSS Masking Module Level 1	CR
mask-border-outset	CSS Masking Module Level 1	CR
mask-border-repeat	CSS Masking Module Level 1	CR
mask-border-slice	CSS Masking Module Level 1	CR
mask-border-source	CSS Masking Module Level 1	CR
mask-border-width	CSS Masking Module Level 1	CR
mask-clip	CSS Masking Module Level 1	CR
mask-composite	CSS Masking Module Level 1	CR
mask-image	CSS Masking Module Level 1	CR
mask-mode	CSS Masking Module Level 1	CR
mask-origin	CSS Masking Module Level 1	CR
mask-position	CSS Masking Module Level 1	CR
mask-repeat	CSS Masking Module Level 1	CR
mask-size	CSS Masking Module Level 1	CR
mask-type	CSS Masking Module Level 1	CR
max-height	CSS 2.1	REC
max-width	CSS 2.1	REC
min-height	CSS 2.1	REC
min-width	CSS 2.1	REC
mix-blend-mode	Compositing and Blending Level 1	CR
object-fit	CSS Images Module Level 3	CR
object-position	CSS Images Module Level 3	CR
opacity	CSS Color Module Level 3	REC
order	CSS Flexible Box Layout Module Level 1	CR
orphans	CSS 2.1	REC
orphans	CSS Fragmentation Module Level 3	CR
outline	CSS 2.1	REC
outline	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
outline-color	CSS 2.1	REC
outline-color	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
outline-offset	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
outline-style	CSS 2.1	REC
outline-style	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
outline-width	CSS 2.1	REC
outline-width	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
overflow	CSS 2.1	REC
overflow-wrap	CSS Text Module Level 3	CR
padding	CSS 2.1	REC
padding	CSS Box Model Module Level 3	CR
padding-bottom	CSS 2.1	REC
padding-bottom	CSS Box Model Module Level 3	CR
padding-left	CSS 2.1	REC
padding-left	CSS Box Model Module Level 3	CR
padding-right	CSS 2.1	REC
padding-right	CSS Box Model Module Level 3	CR
padding-top	CSS 2.1	REC
padding-top	CSS Box Model Module Level 3	CR
page-break-after	CSS 2.1	REC
page-break-before	CSS 2.1	REC
page-break-inside	CSS 2.1	REC
pause	CSS 2.1	REC
pause	CSS Speech Module	CR
pause-after	CSS 2.1	REC
pause-after	CSS Speech Module	CR
pause-before	CSS 2.1	REC
pause-before	CSS Speech Module	CR
pitch	CSS 2.1	REC
pitch-range	CSS 2.1	REC
play-during	CSS 2.1	REC
position	CSS 2.1	REC
quotes	CSS 2.1	REC
resize	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
rest	CSS Speech Module	CR
rest-after	CSS Speech Module	CR
rest-before	CSS Speech Module	CR
richness	CSS 2.1	REC
right	CSS 2.1	REC
scroll-margin	CSS Scroll Snap Module Level 1	CR
scroll-margin-block	CSS Scroll Snap Module Level 1	CR
scroll-margin-block-end	CSS Scroll Snap Module Level 1	CR
scroll-margin-block-start	CSS Scroll Snap Module Level 1	CR
scroll-margin-bottom	CSS Scroll Snap Module Level 1	CR
scroll-margin-inline	CSS Scroll Snap Module Level 1	CR
scroll-margin-inline-end	CSS Scroll Snap Module Level 1	CR
scroll-margin-inline-start	CSS Scroll Snap Module Level 1	CR
scroll-margin-left	CSS Scroll Snap Module Level 1	CR
scroll-margin-right	CSS Scroll Snap Module Level 1	CR
scroll-margin-top	CSS Scroll Snap Module Level 1	CR
scroll-padding	CSS Scroll Snap Module Level 1	CR
scroll-padding-block	CSS Scroll Snap Module Level 1	CR
scroll-padding-block-end	CSS Scroll Snap Module Level 1	CR
scroll-padding-block-start	CSS Scroll Snap Module Level 1	CR
scroll-padding-bottom	CSS Scroll Snap Module Level 1	CR
scroll-padding-inline	CSS Scroll Snap Module Level 1	CR
scroll-padding-inline-end	CSS Scroll Snap Module Level 1	CR
scroll-padding-inline-start	CSS Scroll Snap Module Level 1	CR
scroll-padding-left	CSS Scroll Snap Module Level 1	CR
scroll-padding-right	CSS Scroll Snap Module Level 1	CR
scroll-padding-top	CSS Scroll Snap Module Level 1	CR
scroll-snap-align	CSS Scroll Snap Module Level 1	CR
scroll-snap-stop	CSS Scroll Snap Module Level 1	CR
scroll-snap-type	CSS Scroll Snap Module Level 1	CR
shape-image-threshold	CSS Shapes Module Level 1	CR
shape-margin	CSS Shapes Module Level 1	CR
shape-outside	CSS Shapes Module Level 1	CR
speak	CSS 2.1	REC
speak	CSS Speech Module	CR
speak-as	CSS Speech Module	CR
speak-header	CSS 2.1	REC
speak-numeral	CSS 2.1	REC
speak-punctuation	CSS 2.1	REC
speech-rate	CSS 2.1	REC
stress	CSS 2.1	REC
tab-size	CSS Text Module Level 3	CR
table-layout	CSS 2.1	REC
text-align	CSS 2.1	REC
text-align	CSS Text Module Level 3	CR
text-align-all	CSS Text Module Level 3	CR
text-align-last	CSS Text Module Level 3	CR
text-combine-upright	CSS Writing Modes Level 3	REC
text-combine-upright	CSS Writing Modes Level 4	CR
text-decoration	CSS 2.1	REC
text-decoration	CSS Text Decoration Module Level 3	CR
text-decoration-color	CSS Text Decoration Module Level 3	CR
text-decoration-line	CSS Text Decoration Module Level 3	CR
text-decoration-style	CSS Text Decoration Module Level 3	CR
text-emphasis	CSS Text Decoration Module Level 3	CR
text-emphasis-color	CSS Text Decoration Module Level 3	CR
text-emphasis-position	CSS Text Decoration Module Level 3	CR
text-emphasis-style	CSS Text Decoration Module Level 3	CR
text-indent	CSS 2.1	REC
text-indent	CSS Text Module Level 3	CR
text-justify	CSS Text Module Level 3	CR
text-orientation	CSS Writing Modes Level 3	REC
text-orientation	CSS Writing Modes Level 4	CR
text-overflow	CSS Basic User Interface Module Level 3 (CSS3 UI)	REC
text-shadow	CSS Text Decoration Module Level 3	CR
text-transform	CSS 2.1	REC
text-transform	CSS Text Module Level 3	CR
text-underline-position	CSS Text Decoration Module Level 3	CR
top	CSS 2.1	REC
transform	CSS Transforms Module Level 1	CR
transform-box	CSS Transforms Module Level 1	CR
transform-origin	CSS Transforms Module Level 1	CR
unicode-bidi	CSS 2.1	REC
unicode-bidi	CSS Writing Modes Level 3	REC
unicode-bidi	CSS Writing Modes Level 4	CR
vertical-align	CSS 2.1	REC
visibility	CSS 2.1	REC
voice-balance	CSS Speech Module	CR
voice-duration	CSS Speech Module	CR
voice-family	CSS 2.1	REC
voice-family	CSS Speech Module	CR
voice-pitch	CSS Speech Module	CR
voice-range	CSS Speech Module	CR
voice-rate	CSS Speech Module	CR
voice-stress	CSS Speech Module	CR
voice-volume	CSS Speech Module	CR
volume	CSS 2.1	REC
white-space	CSS 2.1	REC
white-space	CSS Text Module Level 3	CR
widows	CSS 2.1	REC
widows	CSS Fragmentation Module Level 3	CR
width	CSS 2.1	REC
will-change	CSS Will Change Module Level 1	CR
word-break	CSS Text Module Level 3	CR
word-spacing	CSS 2.1	REC
word-spacing	CSS Text Module Level 3	CR
word-wrap	CSS Text Module Level 3	CR
writing-mode	CSS Writing Modes Level 3	REC
writing-mode	CSS Writing Modes Level 4	CR
z-index	CSS 2.1	REC

https://www.w3.org/Style/CSS/all-properties.en.html
Bert Bos, style activity lead
Copyright © 1994–2020 W3C®
*/

//CSS is a Cascading Style Sheet property
type CSS string

// Properties is a map of CSS properties and values
type Properties map[CSS]string

// StyleSheet object is a remote reference to a stylesheet or a map of Properties and Values
// If a URL is set, it will supercede any set properties, and the StyleSheet will produce a link element
// referencing the URL
type StyleSheet struct {
	URL        string
	Properties *Properties
}

const (
	//AlignContent             Defines the alignment between the lines inside a flexible container when the items do not use all available space
	AlignContent = CSS(`align-content`)
	//AlignItems               Defines the alignment for items inside a flexible container
	AlignItems = CSS(`align-items`)
	//AlignSelf                Defines the alignment for selected items inside a flexible container
	AlignSelf = CSS(`align-self`)
	//All                      Resets all properties (except unicode-bidi and direction)
	All = CSS(`all`)
	//Animation                Shorthand for all the animation-* properties
	Animation = CSS(`animation`)
	//AnimationDelay           Defines a delay for the start of an animation
	AnimationDelay = CSS(`animation-delay`)
	//AnimationDirection       Defines whether an animation should be played forwards, backwards or in alternate cycles
	AnimationDirection = CSS(`animation-direction`)
	//AnimationDuration        Defines how long an animation should take to complete one cycle
	AnimationDuration = CSS(`animation-duration`)
	//AnimationFillMode        Defines a style for the element when the animation is not playing (before it starts, after it ends, or both)
	AnimationFillMode = CSS(`animation-fill-mode`)
	//AnimationIterationCount  Defines the number of times an animation should be played
	AnimationIterationCount = CSS(`animation-iteration-count`)
	//AnimationName            Defines a name for the @keyframes animation
	AnimationName = CSS(`animation-name`)
	//AnimationPlayState       Defines whether the animation is running or paused
	AnimationPlayState = CSS(`animation-play-state`)
	//AnimationTimingFunction  Defines the speed curve of an animation
	AnimationTimingFunction = CSS(`animation-timing-function`)
	//BackfaceVisibility       Defines whether  the back face of an element should be visible when facing the user
	BackfaceVisibility = CSS(`backface-visibility`)
	//Background               Shorthand for all the background* properties
	Background = CSS(`background`)
	//BackgroundAttachment     Defines whether a background image scrolls with the rest of the page, or is fixed
	BackgroundAttachment = CSS(`background-attachment`)
	//BackgroundBlendMode      Defines the blending mode of each background layer (color/image)
	BackgroundBlendMode = CSS(`background-blend-mode`)
	//BackgroundClip           Defines how far the background (color or image) should extend within an element
	BackgroundClip = CSS(`background-clip`)
	//BackgroundColor          Defines the background color of an element
	BackgroundColor = CSS(`background-color`)
	//BackgroundImage          Defines one or more background images for an element
	BackgroundImage = CSS(`background-image`)
	//BackgroundOrigin         Defines the origin position of a background image
	BackgroundOrigin = CSS(`background-origin`)
	//BackgroundPosition       Defines the position of a background image
	BackgroundPosition = CSS(`background-position`)
	//BackgroundRepeat         Defines if/how a background image will be repeated
	BackgroundRepeat = CSS(`background-repeat`)
	//BackgroundSize           Defines the size of the background images
	BackgroundSize = CSS(`background-size`)
	//Border                   Shorthand for borderWidth, border-style and border-color
	Border = CSS(`border`)
	//BorderBottom             Shorthand for border-bottom-width, border-bottom-style and border-bottom-color
	BorderBottom = CSS(`border-bottom`)
	//BorderBottomColor        Defines the color of the bottom border
	BorderBottomColor = CSS(`border-bottom-color`)
	//BorderBottomLeftRadius   Defines the radius of the border of the bottom-left corner
	BorderBottomLeftRadius = CSS(`border-bottom-left-radius`)
	//BorderBottomRightRadius  Defines the radius of the border of the bottom-right corner
	BorderBottomRightRadius = CSS(`border-bottom-right-radius`)
	//BorderBottomStyle        Defines the style of the bottom border
	BorderBottomStyle = CSS(`border-bottom-style`)
	//BorderBottomWidth        Defines the width of the bottom border
	BorderBottomWidth = CSS(`border-bottom-width`)
	//BorderCollapse           Defines whether table borders should collapse into a single border or be separated
	BorderCollapse = CSS(`border-collapse`)
	//BorderColor              Defines the color of the four borders
	BorderColor = CSS(`border-color`)
	//BorderImage              Shorthand for all the border-image-* properties
	BorderImage = CSS(`border-image`)
	//BorderImageOutset        Defines the amount by which the border image area extends beyond the border box
	BorderImageOutset = CSS(`border-image-outset`)
	//BorderImageRepeat        Defines whether the border image should be repeated, rounded or stretched
	BorderImageRepeat = CSS(`border-image-repeat`)
	//BorderImageSlice         Defines how to slice the border image
	BorderImageSlice = CSS(`border-image-slice`)
	//BorderImageSource        Defines the path to the image to be used as a border
	BorderImageSource = CSS(`border-image-source`)
	//BorderImageWidth         Defines the width of the border image
	BorderImageWidth = CSS(`border-image-width`)
	//BorderLeft               Shorthand for all the border-left-* properties
	BorderLeft = CSS(`border-left`)
	//BorderLeftColor          Defines the color of the left border
	BorderLeftColor = CSS(`border-left-color`)
	//BorderLeftStyle          Defines the style of the left border
	BorderLeftStyle = CSS(`border-left-style`)
	//BorderLeftWidth          Defines the width of the left border
	BorderLeftWidth = CSS(`border-left-width`)
	//BorderRadius             Shorthand for the four border-*-radius properties
	BorderRadius = CSS(`border-radius`)
	//BorderRight              Shorthand for all the border-right-* properties
	BorderRight = CSS(`border-right`)
	//BorderRightColor         Defines the color of the right border
	BorderRightColor = CSS(`border-right-color`)
	//BorderRightStyle         Defines the style of the right border
	BorderRightStyle = CSS(`border-right-style`)
	//BorderRightWidth         Defines the width of the right border
	BorderRightWidth = CSS(`border-right-width`)
	//BorderSpacing            Defines the distance between the borders of adjacent cells
	BorderSpacing = CSS(`border-spacing`)
	//BorderStyle              Defines the style of the four borders
	BorderStyle = CSS(`border-style`)
	//BorderTop                Shorthand for border-top-width, border-top-style and border-top-color
	BorderTop = CSS(`border-top`)
	//BorderTopColor           Defines the color of the top border
	BorderTopColor = CSS(`border-top-color`)
	//BorderTopLeftRadius      Defines the radius of the border of the top-left corner
	BorderTopLeftRadius = CSS(`border-top-left-radius`)
	//BorderTopRightRadius     Defines the radius of the border of the top-right corner
	BorderTopRightRadius = CSS(`border-top-right-radius`)
	//BorderTopStyle           Defines the style of the top border
	BorderTopStyle = CSS(`border-top-style`)
	//BorderTopWidth           Defines the width of the top border
	BorderTopWidth = CSS(`border-top-width`)
	//BorderWidth              Defines the width of the four borders
	BorderWidth = CSS(`border-width`)
	//Bottom                   Defines the elements position, from the bottom of its parent element
	Bottom = CSS(`bottom`)
	//BoxDecorationBreak       Defines the behavior of the background and border of an element at page-break, or, for in-line elements, at line-break.
	BoxDecorationBreak = CSS(`box-decoration-break`)
	//BoxShadow                Attaches one or more shadows to an element
	BoxShadow = CSS(`box-shadow`)
	//BoxSizing                Defines how the width and height of an element are calculated: should they include padding and borders,
	BoxSizing = CSS(`box-sizing`)
	//BreakAfter               Defines whether  a page-, column-, or region-break should occur after the specified element
	BreakAfter = CSS(`break-after`)
	//BreakBefore              Defines whether  a page-, column-, or region-break should occur before the specified element
	BreakBefore = CSS(`break-before`)
	//BreakInside              Defines whether  a page-, column-, or region-break should occur inside the specified element
	BreakInside = CSS(`break-inside`)
	//CaptionSide              Defines the placement of a table caption
	CaptionSide = CSS(`caption-side`)
	//CaretColor               Defines the color of the cursor (caret) in inputs, textareas, or any element that is editable
	CaretColor = CSS(`caret-color`)
	//Charset                  Defines the character encoding used in the style sheet
	Charset = CSS(`@charset`)
	//Clear                    Defines on which sides of an element floating elements are not allowed to float
	Clear = CSS(`clear`)
	//Clip                     Clips an absolutely positioned element
	Clip = CSS(`clip`)
	//Color                    Defines the color of text
	Color = CSS(`color`)
	//ColumnCount              Defines the number of columns an element should be divided into
	ColumnCount = CSS(`column-count`)
	//ColumnFill               Defines how to fill columns, balanced
	ColumnFill = CSS(`column-fill`)
	//ColumnGap                Defines the gap between the columns
	ColumnGap = CSS(`column-gap`)
	//ColumnRule               Shorthand for all the column-rule-* properties
	ColumnRule = CSS(`column-rule`)
	//ColumnRuleColor          Defines the color of the rule between columns
	ColumnRuleColor = CSS(`column-rule-color`)
	//ColumnRuleStyle          Defines the style of the rule between columns
	ColumnRuleStyle = CSS(`column-rule-style`)
	//ColumnRuleWidth          Defines the width of the rule between columns
	ColumnRuleWidth = CSS(`column-rule-width`)
	//ColumnSpan               Defines how many columns an element should span across
	ColumnSpan = CSS(`column-span`)
	//ColumnWidth              Defines the column width
	ColumnWidth = CSS(`column-width`)
	//Columns                  Shorthand for columnWidth and column-count
	Columns = CSS(`columns`)
	//Content                  Used with the :before and :after pseudoElements, to insert generated content
	Content = CSS(`content`)
	//CounterIncrement         Increases or decreases the value of one or more CSS counters
	CounterIncrement = CSS(`counter-increment`)
	//CounterReset             Creates or resets one or more CSS counters
	CounterReset = CSS(`counter-reset`)
	//Cursor                   Defines the mouse cursor to be displayed when pointing over an element
	Cursor = CSS(`cursor`)
	//Direction                Defines the text direction/writing direction
	Direction = CSS(`direction`)
	//Display                  Defines how a certain HTML element should be displayed
	Display = CSS(`display`)
	//EmptyCells               Defines whether  to display borders and background on empty cells in a table
	EmptyCells = CSS(`empty-cells`)
	//Filter                   Defines effects (e.g. blurring or color shifting) on an element before the element is displayed
	Filter = CSS(`filter`)
	//Flex                     Shorthand for the flexGrow, flex-shrink, and the flex-basis properties
	Flex = CSS(`flex`)
	//FlexBasis                Defines the initial length of a flexible item
	FlexBasis = CSS(`flex-basis`)
	//FlexDirection            Defines the direction of the flexible items
	FlexDirection = CSS(`flex-direction`)
	//FlexFlow                 Shorthand for the flex-direction and the flex-wrap properties
	FlexFlow = CSS(`flex-flow`)
	//FlexGrow                 Defines how much the item will grow relative to the rest
	FlexGrow = CSS(`flex-grow`)
	//FlexShrink               Defines how the item will shrink relative to the rest
	FlexShrink = CSS(`flex-shrink`)
	//FlexWrap                 Defines whether the flexible items should wrap
	FlexWrap = CSS(`flex-wrap`)
	//Float                    Defines whether  a box should float
	Float = CSS(`float`)
	//Font                     Shorthand for the fontStyle, font-variant, font-weight, font-size/line-height, and the font-family properties
	Font = CSS(`font`)
	//FontFace                 A rule that allows websites to download and use fonts other than the "web-safe" fonts
	FontFace = CSS(`@font-face`)
	//FontFamily               Defines the font family for text
	FontFamily = CSS(`font-family`)
	//FontFeatureSettings      Allows control over advanced typographic features in OpenType fonts
	FontFeatureSettings = CSS(`font-feature-settings`)
	//FontFeatureValues        Allows authors to use a common name in font-variant-alternate for feature activated differently in OpenType
	FontFeatureValues = CSS(`@font-feature-values`)
	//FontKerning              Controls the usage of the kerning information (how letters are spaced)
	FontKerning = CSS(`font-kerning`)
	//FontLanguageOverride     Controls the usage of language-specific glyphs in a typeface
	FontLanguageOverride = CSS(`font-language-override`)
	//FontSize                 Defines the font size of text
	FontSize = CSS(`font-size`)
	//FontSizeAdjust           Preserves the readability of text when font fallback occurs
	FontSizeAdjust = CSS(`font-size-adjust`)
	//FontStretch              Selects a normal, condensed, or expanded face from a font family
	FontStretch = CSS(`font-stretch`)
	//FontStyle                Defines the font style for text
	FontStyle = CSS(`font-style`)
	//FontSynthesis            Controls which missing typefaces (bold or italic) may be synthesized by the browser
	FontSynthesis = CSS(`font-synthesis`)
	//FontVariant              Defines whether  a text should be displayed in a small-caps font
	FontVariant = CSS(`font-variant`)
	//FontVariantAlternates    Controls the usage of alternate glyphs associated to alternative names defined in @font-feature-values
	FontVariantAlternates = CSS(`font-variant-alternates`)
	//FontVariantCaps          Controls the usage of alternate glyphs for capital letters
	FontVariantCaps = CSS(`font-variant-caps`)
	//FontVariantEastAsian     Controls the usage of alternate glyphs for East Asian scripts (e.g Japanese and Chinese)
	FontVariantEastAsian = CSS(`font-variant-east-asian`)
	//FontVariantLigatures     Controls which ligatures and contextual forms are used in textual content of the elements it applies to
	FontVariantLigatures = CSS(`font-variant-ligatures`)
	//FontVariantNumeric       Controls the usage of alternate glyphs for numbers, fractions, and ordinal markers
	FontVariantNumeric = CSS(`font-variant-numeric`)
	//FontVariantPosition      Controls the usage of alternate glyphs of smaller size positioned as superscript or subscript regarding the baseline of the font
	FontVariantPosition = CSS(`font-variant-position`)
	//FontWeight               Defines the weight of a font
	FontWeight = CSS(`font-weight`)
	//Grid                     Shorthand for the gridTemplate-rows, grid-template-columns, grid-template-areas, grid-auto-rows, grid-auto-columns, and the grid-auto-flow properties
	Grid = CSS(`grid`)
	//GridArea                 Either specifies a name for the grid item, or this property is a shorthand for the grid-row-start, grid-column-start, grid-row-end, and grid-column-end properties
	GridArea = CSS(`grid-area`)
	//GridAutoColumns          Defines a default column size
	GridAutoColumns = CSS(`grid-auto-columns`)
	//GridAutoFlow             Defines how auto-placed items are inserted in the grid
	GridAutoFlow = CSS(`grid-auto-flow`)
	//GridAutoRows             Defines a default row size
	GridAutoRows = CSS(`grid-auto-rows`)
	//GridColumn               Shorthand for the grid-column-start and the grid-column-end properties
	GridColumn = CSS(`grid-column`)
	//GridColumnEnd            Defines where to end the grid item
	GridColumnEnd = CSS(`grid-column-end`)
	//GridColumnGap            Defines the size of the gap between columns
	GridColumnGap = CSS(`grid-column-gap`)
	//GridColumnStart          Defines where to start the grid item
	GridColumnStart = CSS(`grid-column-start`)
	//GridGap                  Shorthand for the grid-row-gap and grid-column-gap properties
	GridGap = CSS(`grid-gap`)
	//GridRow                  Shorthand for the grid-row-start and the grid-row-end properties
	GridRow = CSS(`grid-row`)
	//GridRowEnd               Defines where to end the grid item
	GridRowEnd = CSS(`grid-row-end`)
	//GridRowGap               Defines the size of the gap between rows
	GridRowGap = CSS(`grid-row-gap`)
	//GridRowStart             Defines where to start the grid item
	GridRowStart = CSS(`grid-row-start`)
	//GridTemplate             Shorthand for the grid-template-rows, grid-template-columns and grid-areas properties
	GridTemplate = CSS(`grid-template`)
	//GridTemplateAreas        Defines how to display columns and rows, using named grid items
	GridTemplateAreas = CSS(`grid-template-areas`)
	//GridTemplateColumns      Defines the size of the columns, and how many columns in a grid layout
	GridTemplateColumns = CSS(`grid-template-columns`)
	//GridTemplateRows         Defines the size of the rows in a grid layout
	GridTemplateRows = CSS(`grid-template-rows`)
	//HangingPunctuation       Defines whether a punctuation character may be placed outside the line box
	HangingPunctuation = CSS(`hanging-punctuation`)
	//Height                   Defines the height of an element
	Height = CSS(`height`)
	//Hyphens                  Defines how to split words to improve the layout of paragraphs
	Hyphens = CSS(`hyphens`)
	//ImageRendering           Gives a hint to the browser about what aspects of an image are most important to preserve when the image is scaled
	ImageRendering = CSS(`image-rendering`)
	//Import                   Allows you to import a style sheet into another style sheet
	Import = CSS(`@import`)
	//Isolation                Defines whether an element must create a new stacking content
	Isolation = CSS(`isolation`)
	//JustifyContent           Defines the alignment between the items inside a flexible container when the items do not use all available space
	JustifyContent = CSS(`justify-content`)
	//Keyframes                Defines the animation code
	Keyframes = CSS(`@keyframes`)
	//Left                     Defines the left position of a positioned element
	Left = CSS(`left`)
	//LetterSpacing            Increases or decreases the space between characters in a text
	LetterSpacing = CSS(`letter-spacing`)
	//LineBreak                Defines how/if to break lines
	LineBreak = CSS(`line-break`)
	//LineHeight               Defines the line height
	LineHeight = CSS(`line-height`)
	//ListStyle                Defines all the properties for a list in one declaration
	ListStyle = CSS(`list-style`)
	//ListStyleImage           Defines an image as the list-item marker
	ListStyleImage = CSS(`list-style-image`)
	//ListStylePosition        Defines the position of the list-item markers (bullet points)
	ListStylePosition = CSS(`list-style-position`)
	//ListStyleType            Defines the type of list-item marker
	ListStyleType = CSS(`list-style-type`)
	//Margin                   Defines all the margin properties in one declaration
	Margin = CSS(`margin`)
	//MarginBottom             Defines the bottom margin of an element
	MarginBottom = CSS(`margin-bottom`)
	//MarginLeft               Defines the left margin of an element
	MarginLeft = CSS(`margin-left`)
	//MarginRight              Defines the right margin of an element
	MarginRight = CSS(`margin-right`)
	//MarginTop                Defines the top margin of an element
	MarginTop = CSS(`margin-top`)
	//Mask                     Hides an element by masking or clipping the image at specific places
	Mask = CSS(`mask`)
	//MaskType                 Defines whether a mask element is used as a luminance or an alpha mask
	MaskType = CSS(`mask-type`)
	//MaxHeight                Defines the maximum height of an element
	MaxHeight = CSS(`max-height`)
	//MaxWidth                 Defines the maximum width of an element
	MaxWidth = CSS(`max-width`)
	//Media                    Defines the style rules for different media types/devices/sizes
	Media = CSS(`@media`)
	//MinHeight                Defines the minimum height of an element
	MinHeight = CSS(`min-height`)
	//MinWidth                 Defines the minimum width of an element
	MinWidth = CSS(`min-width`)
	//MixBlendMode             Defines how an element's content should blend with its direct parent background
	MixBlendMode = CSS(`mix-blend-mode`)
	//ObjectFit                Defines how the contents of a replaced element should be fitted to the box established by its used height and width
	ObjectFit = CSS(`object-fit`)
	//ObjectPosition           Defines the alignment of the replaced element inside its box
	ObjectPosition = CSS(`object-position`)
	//Opacity                  Defines the opacity level for an element
	Opacity = CSS(`opacity`)
	//Order                    Defines the order of the flexible item, relative to the rest
	Order = CSS(`order`)
	//Orphans                  Defines the minimum number of lines that must be left at the bottom of a page when a page break occurs inside an element
	Orphans = CSS(`orphans`)
	//Outline                  Shorthand for the outlineWidth, outline-style, and the outline-color properties
	Outline = CSS(`outline`)
	//OutlineColor             Defines the color of an outline
	OutlineColor = CSS(`outline-color`)
	//OutlineOffset            Offsets an outline, and draws it beyond the border edge
	OutlineOffset = CSS(`outline-offset`)
	//OutlineStyle             Defines the style of an outline
	OutlineStyle = CSS(`outline-style`)
	//OutlineWidth             Defines the width of an outline
	OutlineWidth = CSS(`outline-width`)
	//Overflow                 Defines what happens if content overflows an elements box
	Overflow = CSS(`overflow`)
	//OverflowWrap             Defines whether  the browser may break lines within words in order to prevent overflow (when a string is too long to fit its containing box)
	OverflowWrap = CSS(`overflow-wrap`)
	//OverflowX                Defines whether  to clip the left/right edges of the content, if it overflows the element's content area
	OverflowX = CSS(`overflow-x`)
	//OverflowY                Defines whether  to clip the top/bottom edges of the content, if it overflows the element's content area
	OverflowY = CSS(`overflow-y`)
	//Padding                  Shorthand for all the padding* properties
	Padding = CSS(`padding`)
	//PaddingBottom            Defines the bottom padding of an element
	PaddingBottom = CSS(`padding-bottom`)
	//PaddingLeft              Defines the left padding of an element
	PaddingLeft = CSS(`padding-left`)
	//PaddingRight             Defines the right padding of an element
	PaddingRight = CSS(`padding-right`)
	//PaddingTop               Defines the top padding of an element
	PaddingTop = CSS(`padding-top`)
	//PageBreakAfter           Defines the page-break behavior after an element
	PageBreakAfter = CSS(`page-break-after`)
	//PageBreakBefore          Defines the page-break behavior before an element
	PageBreakBefore = CSS(`page-break-before`)
	//PageBreakInside          Defines the page-break behavior inside an element
	PageBreakInside = CSS(`page-break-inside`)
	//Perspective              Gives a 3DPositioned element some perspective
	Perspective = CSS(`perspective`)
	//PerspectiveOrigin        Defines at which position the user is looking at the 3D-positioned element
	PerspectiveOrigin = CSS(`perspective-origin`)
	//PointerEvents            Defines whether  an element reacts to pointer events
	PointerEvents = CSS(`pointer-events`)
	//Position                 Defines the type of positioning method used for an element (static, relative, absolute or fixed)
	Position = CSS(`position`)
	//Quotes                   Defines the type of quotation marks for embedded quotations
	Quotes = CSS(`quotes`)
	//Resize                   Defines if (and how) an element is resizable by the user
	Resize = CSS(`resize`)
	//Right                    Defines the right position of a positioned element
	Right = CSS(`right`)
	//ScrollBehavior           Defines whether to smoothly animate the scroll position in a scrollable box, instead of a straight jump
	ScrollBehavior = CSS(`scroll-behavior`)
	//TabSize                  Defines the width of a tab character
	TabSize = CSS(`tab-size`)
	//TableLayout              Defines the algorithm used to lay out table cells, rows, and columns
	TableLayout = CSS(`table-layout`)
	//TextAlign                Defines the horizontal alignment of text
	TextAlign = CSS(`text-align`)
	//TextAlignLast            Describes how the last line of a block or a line right before a forced line break is aligned when text-align is "justify"
	TextAlignLast = CSS(`text-align-last`)
	//TextCombineUpright       Defines the combination of multiple characters into the space of a single character
	TextCombineUpright = CSS(`text-combine-upright`)
	//TextDecoration           Defines the decoration added to text
	TextDecoration = CSS(`text-decoration`)
	//TextDecorationColor      Defines the color of the text-decoration
	TextDecorationColor = CSS(`text-decoration-color`)
	//TextDecorationLine       Defines the type of line in a text-decoration
	TextDecorationLine = CSS(`text-decoration-line`)
	//TextDecorationStyle      Defines the style of the line in a text decoration
	TextDecorationStyle = CSS(`text-decoration-style`)
	//TextIndent               Defines the indentation of the first line in a text-block
	TextIndent = CSS(`text-indent`)
	//TextJustify              Defines the justification method used when text-align is "justify"
	TextJustify = CSS(`text-justify`)
	//TextOrientation          Defines the orientation of the text in a line
	TextOrientation = CSS(`text-orientation`)
	//TextOverflow             Defines what should happen when text overflows the containing element
	TextOverflow = CSS(`text-overflow`)
	//TextShadow               Adds shadow to text
	TextShadow = CSS(`text-shadow`)
	//TextTransform            Controls the capitalization of text
	TextTransform = CSS(`text-transform`)
	//TextUnderlinePosition    Defines the position of the underline which is set using the text-decoration property
	TextUnderlinePosition = CSS(`text-underline-position`)
	//Top                      Defines the top position of a positioned element
	Top = CSS(`top`)
	//Transform                Applies a 2D or 3D transformation to an element
	Transform = CSS(`transform`)
	//TransformOrigin          Allows you to change the position on transformed elements
	TransformOrigin = CSS(`transform-origin`)
	//TransformStyle           Defines how nested elements are rendered in 3D space
	TransformStyle = CSS(`transform-style`)
	//Transition               Shorthand for all the transition* properties
	Transition = CSS(`transition`)
	//TransitionDelay          Defines when the transition effect will start
	TransitionDelay = CSS(`transition-delay`)
	//TransitionDuration       Defines how many seconds or milliseconds a transition effect takes to complete
	TransitionDuration = CSS(`transition-duration`)
	//TransitionProperty       Defines the name of the CSS property the transition effect is for
	TransitionProperty = CSS(`transition-property`)
	//TransitionTimingFunction Defines the speed curve of the transition effect
	TransitionTimingFunction = CSS(`transition-timing-function`)
	//UnicodeBidi              Used together with the direction property to set or return whether the text should be overridden to support multiple languages in the same document
	UnicodeBidi = CSS(`unicode-bidi`)
	//UserSelect               Defines whether the text of an element can be selected
	UserSelect = CSS(`user-select`)
	//VerticalAlign            Defines the vertical alignment of an element
	VerticalAlign = CSS(`vertical-align`)
	//Visibility               Defines whether  an element is visible
	Visibility = CSS(`visibility`)
	//WhiteSpace               Defines how white-space inside an element is handled
	WhiteSpace = CSS(`white-space`)
	//Widows                   Defines the minimum number of lines that must be left at the top of a page when a page break occurs inside an element
	Widows = CSS(`widows`)
	//Width                    Defines the width of an element
	Width = CSS(`width`)
	//WordBreak                Defines how words should break when reaching the end of a line
	WordBreak = CSS(`word-break`)
	//WordSpacing              Increases or decreases the space between words in a text
	WordSpacing = CSS(`word-spacing`)
	//WordWrap                 Allows long, unbreakable words to be broken and wrap to the next line
	WordWrap = CSS(`word-wrap`)
	//WritingMode              Defines whether lines of text are laid out horizontally or vertically
	WritingMode = CSS(`writing-mode`)
	//ZIndex                   Defines the stack order of a positioned element
	ZIndex = CSS(`z-index`)
)

func (ps *Properties) String() string {
	css := ""
	keys := make([]string, 0, len(*ps))
	for k := range *ps {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := (*ps)[CSS(k)]
		css += fmt.Sprintf("%s:%s;", k, v)
	}
	return css
}

func (ss *StyleSheet) String() string {
	if ss.URL != "" {
		return fmt.Sprintf(`<link rel="stylesheet" href="%s">`, ss.URL)
	}
	return fmt.Sprintf("%s", ss.Properties)
}

// AddProperty adds a property and its value to a stylesheet
func (ss *StyleSheet) AddProperty(property CSS, value string) {
	if ss.Properties == nil {
		ss.Properties = &Properties{}
	}
	(*ss.Properties)[property] = value
}

// NewStyleSheet generates a new CSS Style Sheet object
func NewStyleSheet() *StyleSheet {
	return &StyleSheet{
		URL:        "",
		Properties: &Properties{},
	}
}
