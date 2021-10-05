---
# try also 'default' to start simple
theme: default
# random image from a curated Unsplash collection by Anthony
# like them? see https://unsplash.com/collections/94734566/slidev
# background: ./public/vid/01.mp4
layout: cover
# apply any windi css classes to the current slide
class: 'text-center'
# https://sli.dev/custom/highlighters.html
highlighter: shiki
# show line numbers in code blocks
lineNumbers: false
# some information about the slides, markdown enabled
info: |
  ## Slidev Starter Template
  Presentation slides for developers.

  Learn more at [Sli.dev](https://sli.dev)
# persist drawings in exports and build
drawings:
  persist: false
---

<h1 style="color:black">Concurrency Patterns</h1>

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/background_yellow.mp4" type="video/mp4">
</video>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/concurrent_world.mp4" type="video/mp4">
</video>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/more_people.mp4" type="video/mp4">
</video>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/elevators.mp4" type="video/mp4">
</video>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/pegeons.mp4" type="video/mp4">
</video>

---
layout: image
image: /img/one_way.jpg
---

---
layout: statement
---

# CONCURRENCY<br>

<h1 style="color:red;">IS NOT</h1>

# PARALLELISM

---
layout: image
image: /img/system_thread_main.png
---

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/airport_conc.mp4" type="video/mp4">
</video>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/airport_conc2.mp4" type="video/mp4">
</video>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/airport_par.mp4" type="video/mp4">
</video>

<div style="position: absolute; bottom: 0px; right: 0px; background-color: black; min-height: 50px; min-width: 400px"></div>

---
layout: image
image: /img/moore.png
---

---
layout: image
image: /img/exponential.png
---

---
layout: image
image: /img/ibm.jpg
---

---
layout: image
image: /img/cloud_fix.png
---

---
src: ./slides/classification_system.md
---

---
layout: image
image: /img/process_full.png
---

<h2 style="color: black; font-size: 40px; text-align: center; margin-top: 60px;">Threading is better than Multiprocessing</h2>

<br>

<div v-click style="color: black; text-align: center;">... for single-domain applications</div>

---
src: ./slides/multiprocessing.md
---

---
layout: image
image: /img/system_thread_main.png
---

<h2 style="margin: 0 0 0 30%; color: black; font-size: 48px;">Threading</h2>

---
src: ./slides/threading_cpp.md
---

<!-- --- -->
<!-- layout: image -->
<!-- image: /img/ideal_threads_ex2.png -->
<!-- --- -->

---
src: ./slides/classification_noasync.md
---

---
layout: statement
---

## But what if I want <b><span class="text-red-500">a lot</span></b> of threads?

---
src: ./slides/goroutines.md
---

---
layout: center
---

You are given a task to add 1 to each integer in a list.


TODO: you are here. Probably time to start with problems

---
layout: image
image: '/img/birds_sharing.jpg'
---

---
src: ./slides/cpp_erlang_comparison.md
---

---
src: ./slides/sharing_problems.md
---

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/racecars.mp4" type="video/mp4">
</video>

---
src: ./slides/race_condition_formal.md
---

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/train_single_1.mp4" type="video/mp4">
</video>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/train_collision.mp4" type="video/mp4">
</video>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/mov_train_fun.mp4" type="video/mp4">
</video>


---
src: ./slides/race_condition_go.md
---

---
src: ./slides/semaphore.md
---

---
src: ./slides/mutex.md
---

---
src: ./slides/mutex_go.md
---

---
src: ./slides/deadlock_pdd.md
---

---
layout: image
image: /img/deadlock.jpg
---

---
src: ./slides/deadlock_go.md
---

---
layout: statement
---

# Livelock

---
src: ./slides/classification.md
---

---

## Actor model and Process calculus (coroutines)

...

---
layout: image
image: /img/idris.jpg
---

<h2 style="color: black; font-size: 42px;">Actor model</h2>

---
layout: image
image: /img/erlang_scala.png
---

---
src: ./slides/actor_erlang.md
---

---
src: ./slides/actor_scala.md
---

---

# Components


<Counter :count="10" m="t-4" />


---
preload: false
---

# Animations

Animations are powered by [@vueuse/motion](https://motion.vueuse.org/).

```html
<div
  v-motion
  :initial="{ x: -80 }"
  :enter="{ x: 0 }">
  Slidev
</div>
```

<div class="w-60 relative mt-6">
  <div class="relative w-40 h-40">
    <img
      v-motion
      :initial="{ x: 800, y: -100, scale: 1.5, rotate: -50 }"
      :enter="final"
      class="absolute top-0 left-0 right-0 bottom-0"
      src="https://sli.dev/logo-square.png"
    />
    <img
      v-motion
      :initial="{ y: 500, x: -100, scale: 2 }"
      :enter="final"
      class="absolute top-0 left-0 right-0 bottom-0"
      src="https://sli.dev/logo-circle.png"
    />
    <img
      v-motion
      :initial="{ x: 600, y: 400, scale: 2, rotate: 100 }"
      :enter="final"
      class="absolute top-0 left-0 right-0 bottom-0"
      src="https://sli.dev/logo-triangle.png"
    />
  </div>

  <div
    class="text-5xl absolute top-14 left-40 text-[#2B90B6] -z-1"
    v-motion
    :initial="{ x: -80, opacity: 0}"
    :enter="{ x: 0, opacity: 1, transition: { delay: 2000, duration: 1000 } }">
    Slidev
  </div>
</div>

<!-- vue script setup scripts can be directly used in markdown, and will only affects current page -->
<script setup lang="ts">
const final = {
  x: 0,
  y: 0,
  rotate: 0,
  scale: 1,
  transition: {
    type: 'spring',
    damping: 10,
    stiffness: 20,
    mass: 2
  }
}
</script>

<div
  v-motion
  :initial="{ x:35, y: 40, opacity: 0}"
  :enter="{ y: 0, opacity: 1, transition: { delay: 3500 } }">

[Learn More](https://sli.dev/guide/animations.html#motion)

</div>
