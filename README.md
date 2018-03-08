# Pin Board Display

## Functionalities

### Editor

Editor is the default screen. Every new editor will have a unique identified at
the end for identifying a new post.

URLs:

* https://editor.local/edit/A087E305-F43B-4B01-97E8-A7222B43E52C
* https://editor.local/view/A087E305-F43B-4B01-97E8-A7222B43E52C

### Kiosk

Kiosk mode is for displaying the post. There can be more than one post in a
kiosk page.

* https://kiosk.local/A087E305-F43B-4B01-97E8-A7222B43E52C

It should have a functionality of refreshing posts/page automatically at every
interval.

## Architecture

### Storage

* Long term storage MongoDB

### Server

* Golang Gorilla Mux
