# JPEG So Small!
JPEG So Small! is a simple application to adjust JPEG quality
## Usage
select input directory and output directory. And it will compress all the jpeg inside input directory

## Setup

jpeg-so-small is written using [Wails](https://wails.io/) with Svelte template

### Requirements
- Go (>= v1.23)
- pnpm (or npm/yarn, edit configuration in `wails.json`)
- upx (optional for compression)
- Wails ([installation](https://wails.io/docs/gettingstarted/installation))

### Live Development

```bash
wails dev
```
### Build Production
```bash
wails build -upx
```
