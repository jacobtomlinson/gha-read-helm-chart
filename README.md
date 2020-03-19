# Read Helm Chart

[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-Read%20Helm%20Chart-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/jacobtomlinson/gha-read-helm-chart)
[![Actions Status](https://github.com/jacobtomlinson/gha-read-helm-chart/workflows/Build/badge.svg)](https://github.com/jacobtomlinson/gha-read-helm-chart/actions)
[![Actions Status](https://github.com/jacobtomlinson/gha-read-helm-chart/workflows/Integration%20Test/badge.svg)](https://github.com/jacobtomlinson/gha-read-helm-chart/actions)

This action will read a Helm Chart's `Chart.yaml` file and expose the values as outputs.

## Usage

Describe how to use your action here.

### Example workflow

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Read Helm Chart
        uses: jacobtomlinson/gha-read-helm-chart@master
        with:
          path: path/to/chart
```

### Inputs

| Input  | Description                                                  |
| ------ | ------------------------------------------------------------ |
| `path` | Path to the helm chart relative to the root of your project. |

### Outputs

| Output                          | Description                                                     |
| ------------------------------- | --------------------------------------------------------------- |
| `apiVersion`                    | The chart API version _(always set)_                            |
| `name`                          | The name of the chart _(always set)_                            |
| `version`                       | A SemVer 2 version _(always set)_                               |
| `kubeVersion`                   | A SemVer range of compatible Kubernetes versions _(optional)_   |
| `description`                   | A single-sentence description of this project _(optional)_      |
| `type`                          | It is the type of chart _(optional)_                            |
| `keywords`                      | A list of keywords about this project _(optional)_              |
| `home`                          | The URL of this project's home page _(optional)_                |
| `sources`                       | A list of URLs to source code for this project _(optional)_     |
| `depenencies_{name}_version`    | Version of dependency `{name}` _(optional)_                     |
| `depenencies_{name}_repository` | Repository URL of dependency `{name}` _(optional)_              |
| `repository`                    | The repository URL _(optional)_                                 |
| `icon`                          | A URL to an SVG or PNG image to be used as an icon _(optional)_ |
| `appVersion`                    | The version of the app that this contains _(optional)_          |
| `deprecated`                    | Whether this chart is deprecated _(optional)_                   |

## Examples

### Using outputs

```yaml
steps:
- uses: actions/checkout@master
- name: Read Helm Chart
  id: chart
  uses: jacobtomlinson/gha-read-helm-chart@master
  with:
    path: path/to/chart
- name: Print outputs
    run: |
    echo "Name - ${{ steps.chart.outputs.name }}"
    echo "Version - ${{ steps.chart.outputs.version }}"
    echo "App Version - ${{ steps.chart.outputs.appVersion }}"
```
