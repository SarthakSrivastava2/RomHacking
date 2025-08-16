How to edit sprites:
Step one:  Open your narc.
	1a: For Diamond or Pearl:
		i.) Check the 'Diamond/Pearl' check box
		ii.) Click the 'Open narc' button, select "Open narc..." from the menu, or use the keyboard shortcut Ctrl+O
		iii.) Find "pokegra.narc" and hit Open
	1b: For Platinum:
		i.) Make sure the 'Diamond/Pearl' check box is not checked
		ii.) Click the 'Open narc' button, select "Open narc..." from the menu, or use the keyboard shortcut Ctrl+O
		iii.) Find "pl_pokegra.narc" and hit Open
	1c: For Heart Gold or Soul Silver:
		i.) Make sure the 'Diamond/Pearl' check box is not checked
		ii.) Click the 'Open narc' button, select "Open narc..." from the menu, or use the keyboard shortcut Ctrl+O
		iii.) Change the filter from "*.narc" to "*.*"
		iv.) Find "4" and hit Open
Step two:  Using the drop-down box at the top of the form, find the entry you are looking for.
Step three:  Edit the sprites. (Note: If any dialog boxes appear, consult the "Dialog Boxes" section)
	3a: To change a single sprite:
		i.) Click on the picturebox you want to replace
		ii.) Select the PNG file you wish to use
	3b: To create a full set of images from a 256x64 sprite sheet:
		i.) Make sure the sprite sheet is in the order "Normal Front, Shiny Front, Normal Back, Shiny Back"
		ii.) Click the 'Load Sprite Sheet' button, select "Load Sprite Sheet" from the menu, or use the keyboard shortcut Ctrl+I
		iii.) Select the PNG to be used
	3c: To create a shiny palette from a single shiny image:
		i.) Click on the 'shiny' picture box that matches the shiny image you want to use
		ii.) Select the PNG file you wish to use
	3d: To create a shiny palette from two matching images:
		i.) Click on the "Create Shiny Palette" button
		ii.) Select the PNG file to use as a base
		iii.) Select the PNG file that contains the shiny image
	3e: To create a custom image set from 2-5 images:
		i.) Click the 'Load Sprite Set' button, select "Load Sprite Set" from the menu, or use the keyboard shortcut Ctrl+L
		ii.) Select the images you wish to use in their apropriate text boxes, using the loading buttons
		iii.) If selecting an image to use for the shiny palette, slect the radio button for the selected image that it matches.  If none or unknown, 					select 'Other'.
		iv.) If "Autofill" is selected, the program will use the same backsprite for both male and female if only one is selected.  The same is true of 				frontsprites.  Also, the normal palette will be used for the shiny palette if no shiny image is selected.
		v.) If "Match Current Palette" is selected, the program will attempt to match the loaded images to the current palette(except the shiny image).
Step four:  Save your work.
	4: Select "Write to narc..." from the menu or use the keyboard shortcut Ctrl+W

How to extract sprites from the narc:
Step one:  Follow steps one and two from the "How to edit sprites" section
Step two:  Save the sprites you want.
	2a: To save a single sprite:
		i.) Using the drop-down box at the bottom of the form, find the image you want to save
		ii.) Click on the "Save PNG" button.
		iii.) Enter the name you want to save as
		iv.) Hit the "Save" button
	2b: To save an entire set of images:
		i.) Select "Save Sprite Set" from the menu, or use the keyboard shortcut Ctrl+S
		ii.) Enter the filename you want to use
		iii.) NOTE: the program will append a suffix to the filename of each image(FBack, MBack, FFront, MFront, and Shiny)
		iv.) When you are satisfied, hit the "Save" Button.

OPTIONS:
	'Fix Non-standard Colors': When loading an image, automatically fix colors that will not save correctly.
	'Convert Wrong Format Images': When loading an image, attempts to convert images that do not have the correct color format.
	'Allow Shrinking of Expanded Images': When attenpting to load an image that is too large, will offer to shrink images of appropriate dimensions.

Dialog Boxes:
"Image is Empty" -> You tried to save a blank image
"The sprite sheet should be 256x64." -> The sprite sheet is improperly formatted
"<filename> is not 8bpp Indexed!  Attempt conversion?" -> The image you tried to load is not the correct color format.  If you select 'Yes,' the program will attempt to convert the image to the correct format.  If 'No,' the operation will be cancelled.  
"Conversion failed." -> The image you loaded is not a color format that the program can convert, or it contains more than 256 colors
"<filename> size not recognized. Use Canvas Splitter?" -> The image you loaded is not one of the three sizes the program recognizes: 64x64, 80x80, or 160x80.  If you select 'Yes,' the program will open the canvas splitting tool.  If 'No,' the operation will be cancelled.  The Canvas Splitter will ask you for a size recommendation for splitting the image, and then provide you with a number of tiles of that size.  Click the one that you wish to use, or 'Cancel' if none are correct.
"Image's palette contains more than sixteen colors.  Attempt to shrink?" -> The palette of the image you loaded contains too many colors to be saved to the narc
file.  If you select 'Yes,' the program will attempt to cull unused colors from the palette and recheck.  If 'No,' the warning will be ignored.  WARNING: attempting to save an image with to many colors to the narc will likely end in an image that looks nothing like the image you loaded.
"Palette still too large.  Image will not save correctly." -> Palette shrinking did not reduce the number of colors to 16 or less.
"MakeImage Failed" -> The file you opened with "Open narc" is improperly formatted or corrupted.
"Image's palette does not match the current palette.  Use PaletteMatch?" -> Replacing the current palette with the loaded image's palette will cause color errors.  If you select 'Yes,' the program will attempt to create a new palette that contains the colors from the new image as well as the current palette.  If the images are too different, the new palette will be too large to save.  If 'No' the program will replace the current palette with no changes.  