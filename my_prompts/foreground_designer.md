## INSTRUCTIONS
- Help to convert the `USER REQUIREMENT`, `STYLE REQUIREMENT` and `HTML COMPONENT REQUIREMENT` into a JSON structured output.
- `USER REQUIREMENT` is the user requirement, it's like a design direction for banner ad.
- `STYLE REQUIREMENT` is the style requirement, it's like the font, color, size, etc. When you are generating the JSON structured output, you should use the font, color, size, etc from `STYLE REQUIREMENT`.
- `HTML COMPONENT REQUIREMENT` is the HTML component requirement, it's like the image, text, button, etc. When you are generating the JSON structured output, you should use the image, text, button, etc from `HTML COMPONENT REQUIREMENT`.
- **`SAFE ZONE` is a polygon CSS property (clip-path: polygon). All visual elements (text, logo, shape, button, etc.) MUST be entirely placed within this Polygon. Elements are NOT allowed to extend outside of the `SAFE ZONE`**.

### OUTPUT
Always document your layout rationale and ensure all positions are precisely specified in pixels or percentages.


Below are the examples for major layout styles:
```json
{{
    "image": [{{
        "id": "logo",
        "position": {{
            "left": 20,
            "top": 20
        }},
        "relative_position": {{
            "reference": "banner-top-left",
            "offset_x": 20,
            "offset_y": 20
        }},
        "width": 95,
        "height": 60,
        "additional_instructions": "<Additional instructions for the current logo. If above properties can't describe well the design>"
    }}],
    "background_image": {{
        "id": "background_image",
        "width": 1080,
        "height": 1080
    }},
    "text": [{{
        "id": "headline",
        "content": "Hello World",
        "horizontal_alignment": "left",
        "position": {{
            "left": 20,
            "top": 20
        }},
        "container_padding": {{
            "padding_left": 20,
            "padding_top": 20,
            "padding_right": 20,
            "padding_bottom": 20
        }},
        "width": 260,
        "height": 45,
        "font_family": "Arial",
        "font_size": 36,
        "font_weight": "bold",
        "font_color": "#000000",
        "line-height": "30px",
        "background_color": "#FFFFFF", # Don't fill if not necessary, can leave empty
        "additional_instructions": "<Additional instructions for the current headline. If above properties can't describe well the design>"
    }},{{
        "id": "sub_headline",
        "content": "Sub headline 2",
        "horizontal_alignment": "center",
        "position": {{
            "left": 20,
            "top": 20
        }},
        "container_padding": {{
            "padding_left": 20,
            "padding_top": 20,
            "padding_right": 20,
            "padding_bottom": 20
        }},
        "width": 260,
        "height": 45,
        "font_family": "Arial",
        "font_size": 36,
        "font_weight": "bold",
        "font_color": "#000000",
        "line-height": "30px",
        "background_color": "#FFFFFF", # Don't fill if not necessary, can leave empty
        "additional_instructions": "<Additional instructions for the current sub_headline. If above properties can't describe well the design>"
    }},
    ],
    "additional_instructions": "Additional instructions for the general layout."
}}
```

`USER REQUIREMENT`
```
{user_requirement}
```

`HTML COMPONENT REQUIREMENT`
```
{html_component_requirement}
```

`STYLE REQUIREMENT`
```
{style_requirement}
```

`SAFE ZONE`
```
{safe_zone}
```

Analyze those inputs `USER REQUIREMENT`, `HTML COMPONENT REQUIREMENT`, `STYLE REQUIREMENT` and `SAFE ZONE` with dimensions ({width}x{height}px) as required and take the following into consideration.

Provide complete specifications following the schema exactly. The output should follow valid JSON format. 

REMEMBER:
- All visual elements MUST be entirely contained within the SAFE ZONE.
- Keep all the id in the HTML COMPONENT REQUIREMENT the same. Don't add extra id or change the id.
- The font used must totally from `STYLE REQUIREMENT`. DON'T USE ANY OTHER FONT.
- Pay attention to the relative position and spacing each element.