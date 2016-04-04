App structure

Root project is a codrova project

- app
- hooks
- platform
- plugins
- www (similar to dist)
- other node and bower dependencies


- gulp builds
  - dev build   - will copy the entire html and js resources to www for dev debugging
  - prod build  - will vulcanize, crisp, min to www folder
