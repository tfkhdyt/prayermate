# PrayerMate

PrayerMate is a CLI-based application to remind users of prayer times.

## Features

- Display a list of available locations
- Set the default location
- Display the currently selected location
- Display today's prayer schedule
- Display a push notification when it's time to pray.

## Built with

- [Go](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra)
- [Viper](https://github.com/spf13/viper)
- [Tablewriter](https://github.com/olekukonko/tablewriter)
- [Beeep](https://github.com/gen2brain/beeep)

## Installation

### AUR

```bash
yay -S prayermate-id
```

### Build from source

```bash
sudo make install
```

## Usage

### Get list of available locations

```bash
prayermate location list
```

### Set default location

```bash
prayermate location set [location id]
```

### Show currently selected location

```bash
prayermate location current
```

### Show today's prayer schedule

```bash
prayermate show
```

### Display push notification when it's time to pray.

```bash
prayermate notify
```
