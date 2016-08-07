- Introduction
 - [overview ](#overview)
 - setup guide
 - getting started

- generator
 - configuration
 - templating
 - themes

# Introduction

## overview
wuxia is a potable static web solution. Wuxia allows you to create, build and
deploy static websites. I forgot to mention that, wuxia can also host mustiple
static websites too.

The main motivation behind this is to open up ways for people in my country to
write technical documents  i.e books, articles, etc and also have access to the
documents. This can fit to different use cases depending on who is using it.

Wuxia is made up of components that when combined together forms a powerful
solution for the static web fans. It ships with a static site generator, static
site server, a `git push` based static site builder and a multi domain static
site host all in a single binary.

I built this for personal use, but if you find it useful please do the
followwing, [star the project](https://github.com/gernest/wuxia) , [fork the project](https://github.com/gernest/wuxia), [follow the author](https://github.com/gernest)
and in case you encounter any bugs then [open an issue](https://github.com/gernest/wuxia/issues).

# Generator
Wuxia ships with a static site generator. The implementation of generator is as
a package which means it can be used independently  with another project.

### Why another static generator?
This question is inevitable, there is already a ton load of static site
generators out there, so why did I decide we need a room for one more?

- Seamless integration with the current system. Since the system is written in
  Go, It was obvious using a Go based static site generator would be a good
  idea except for the fact that most static site generators are applications and
  not libraries. Wuxia generator is a library that you can integrate into a
  bigger system.

- Metrics: I needed a way not only to have benchmarks  but all the metrics that
  one can gather about the static content generation process.

- Ease to extend with plugins: I wanted a generator that can be extended in
  functionality and not a black box. This throws many Go options out of the
  window. The current implementation uses javascript as a scripting language.
  This opens up a wide array of good stuffs to the static web generation
  experience.

The generator works by passing through three phases.

- The Configuration phase. This is the first phase of the generator. The project
  must have a configuration file. The name of the configuration file can either
  be `config.json` which is in json format, `config.yml` which is in yaml format
  and `config.toml` which is in the toml format. The configuration file must be
  at the root of the project.

  Something to note is, all operations occurs on the root of the project which
  is to be generated. Since wuxia support scripting and allows reading and
  writing of files at the code generation process, it is safe to guarantee that
  access to any file outside of the project will not be allowed.

- The initialization phase. Anything that is needed by the generator to work
  properly for the project happens here. The user can define a custom script
  which will have access to the Generator settings and other important stuffs.
  The script on path `scripts/init/index.js` is evaluated.

- The planning phase. Planning of the execution is done, and if there is any
  custom made plan script in the path `scripts/plan/index.js` will be evaluated
  and used.

- The Execution phase. Executes the plan and exits the process.
