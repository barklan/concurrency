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
layout: image
image: /img/concurrency_ru.png
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

<div v-click style="color: #484893; text-align: center; font-size: 14px;">... for single-domain applications in most cases, terms and conditions apply. Side effects may include nausea and sudden death. Consult your doctor before use.</div>

---
src: ./slides/multiprocessing.md
---

---
layout: image
image: /img/chrome.png
---

<span style="color: black; font-size: ">..but,</span>

---
src: ./slides/gil.md
---

---
layout: image
image: /img/system_thread_main.png
---

<h2 style="margin: 0 0 0 30%; color: black; font-size: 48px;">Threading</h2>

---
src: ./slides/threading_cpp.md
---

---
layout: statement
---

## But what if I want <b><span class="text-red-500">a lot</span></b> of threads?

---
layout: image
image: /img/ideal_threads_ex2.png
---

---
src: ./slides/classification_noasync.md
---

---
src: ./slides/goroutines.md
---

---
layout: statement
---

## But <b><span class="text-red-500">why?</span></b>

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/communication.mp4" type="video/mp4">
</video>

<div style="position: absolute; top: 0px; right: 0px; background-color: black; min-height: 60px; min-width: 100%"></div>

<div style="position: absolute; bottom: 0px; right: 0px; background-color: black; min-height: 60px; min-width: 100%"></div>

---
layout: image
image: '/img/birds_sharing.jpg'
---

---
layout: statement
---

<h2><span class="text-red-500">Race condition.</span></h2>

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

<div style="position: absolute; top: 0px; right: 0px; background-color: black; min-height: 50px; min-width: 100%"></div>

<div style="position: absolute; bottom: 0px; right: 0px; background-color: black; min-height: 45px; min-width: 100%"></div>

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
layout: image
image: /img/semaphore.jpg
---

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/breaking_bad.mp4" type="video/mp4">
</video>

---
src: ./slides/mutex_go.md
---

---
layout: statement
---

<h2>Race condition. <span class="text-red-500">Deadlock.</span></h2>

---
src: ./slides/deadlock_pdd.md
---

---
src: ./slides/deadlock_pdd_go.md
---

---
layout: image
image: /img/deadlock.jpg
---

---
layout: statement
---

<h2>Race condition. Deadlock. <span class="text-red-500">Starvation.</span></h2>

---
layout: statement
---

<h2>Race condition. Deadlock. Starvation. <span class="text-red-500">Livelock.</span></h2>

---
layout: image
image: /img/problems.jpg
---

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/kid_and_math.mp4" type="video/mp4">
</video>

---
src: ./slides/no_lock_communication.md
---

---
layout: image
image: /img/idris_compose.png
---

---
layout: image
image: /img/elevator_compose.png
---

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
layout: image
image: /img/csp_compose.png
---

---
layout: image
image: /img/csp_sum.png
---

---
src: ./slides/go_chan.md
---

---
src: ./slides/race_condition_go_chan.md
---

---
layout: center
---

"generators", "annjoe", "fanin"

---
src: ./slides/classification.md
---

---
layout: image
image: /img/node.png
---

---
src: ./slides/event_loop_intro.md
---

---
preload: false
---

<video autoplay muted loop class="backgroundVideo">
  <source src="/vid/roller.mp4" type="video/mp4">
</video>

<h1 style="border-radius: 10px; padding: 30px; background-color: white; color:black; right: 100px; position: fixed; font-size: 50px">Event Loop</h1>

---
layout: image
image: /img/libuv_wide.png
---

---
src: ./slides/node_uvloop.md
---

---
layout: image
image: /img/event_loop_ticks.png
---

---
layout: image
image: /img/event_loop_monitor1.png
---

---
layout: center
---

Cluster module...

---
src: ./slides/callback_hell_ground.md
---

---
src: ./slides/callback_hell.md
---

---
src: ./slides/promises.md
---

---
layout: center
---

`Promise.race()`, `Promise.all()`?

---
src: ./slides/promise_chain.md
---

---
src: ./slides/async_await1.md
---

---
src: ./slides/async_await_comparison.md
---
