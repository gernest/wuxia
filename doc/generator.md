# wu gen

`gen` is the command for the static website generator that is shippd with wuxia.
The generator is easy to use and extend.

## Overview of features of the generator

- Convention plus configuration. It comes with sane defaults, sane conventions
  of the way the project should be structured. On to p of this, it gives full
  freedom if the user is advanced enough to configure it in way that it suits
  his/her needs.

- Themes : Support customizing of project's look and feel. The default template
  language is is Go. Effort is underway to add popular javascript templating
  engines.

- Javascript Plugins: Write plugins in javascript. You can tap almost every
  aspect of the generation process through javascript. The generator comes with
  a simple module loader `require` which you can use to break down your modules
  into separate files.

- Metrics: The generator is wired with exhaustive metrics. You can have full
  understanding of the resources used and perfomrance of rendering your project.


## The structure of a wuxia static website project

```
├── scripts
│   ├── init
│   │   └── index.js
│   ├── plan
│   │   └── index.js
│   └── plugin
│       └── index.js
├── src
│   ├── front_matter.md
│   └── home.md
├── templates
│   └── index.html
├── LICENCE
├── README.md
└── config.json
```
This is supposed to be the default project layout. It meansthat it is possible
for projects to decide how to structue the project, only one thing is imortant
that is the configuration file that sits aat the root of the project.
