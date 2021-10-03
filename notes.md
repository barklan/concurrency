# Concurrency Patterns

## What is concurrency and parallelism?

In computer science, concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or at the same time simultaneously partial order, without affecting the final outcome. This allows for parallel execution of the concurrent units, which can significantly improve overall speed of the execution in multi-processor and multi-core systems. In more technical terms, concurrency refers to the decomposability of a program, algorithm, or problem into order-independent or partially-ordered components or units of computation.

Примечание — В русскоязычной литературе нередко путаются термины «параллелизм» и «конкурентность». Оба термина означают одновременность процессов, но первый — на физическом уровне (параллельное исполнение нескольких процессов, нацеленное только на повышение скорости исполнения за счёт использования соответствующей аппаратной поддержки), а второй — на логическом (парадигма проектирования систем, идентифицирующая процессы как независимые, что в том числе позволяет их исполнять физически параллельно, но в первую очередь нацелено на упрощение написания многопоточных программ и повышение их устойчивости).

## Весь мир конкурентный

куча слайдов - дорога поменьше -> дорога побольше -> еще больше

Вообще концепты пошли изначально когда нужно было на одном пути управлять несколькими поездами (в данном случае распределенный ресурс - единый путь, агенты - поезда)


## But why?

Через сорок лет после публикации закона Мура продолжающееся возрастание производительности микросхем происходит благодаря методам локального и глобального массового параллелизма. Локальный параллелизм задействован в новых микросхемах для 64-разрядных многоядерных микропроцессоров, в мульти-чиповых модулях и высокопроизводительных системах связи. Глобальный параллелизм в настоящее время задействован в новом оборудовании для проводной и беспроводной широкополосной пакетной коммутации сообщений. Ёмкость хранения за счёт как локального, так и глобального параллелизма, растёт в геометрической прогрессии.

Here moore's law and exponential

**Параллелизм и закон Мура**
В последнее время, чтобы получить возможность задействовать на практике ту дополнительную вычислительную мощность, которую предсказывает закон Мура, стало необходимо задействовать параллельные вычисления. На протяжении многих лет производители процессоров постоянно увеличивали тактовую частоту и параллелизм на уровне инструкций, так что на новых процессорах старые однопоточные приложения исполнялись быстрее без каких-либо изменений в программном коде.[8] Примерно с середины десятилетия 2000-х годов по разным причинам производители процессоров предпочитают многоядерные архитектуры, и для получения всей выгоды от возросшей производительности ЦП программы должны переписываться в соответствующей манере. Однако не каждый алгоритм поддается распараллеливанию, определяя, таким образом, фундаментальный предел эффективности решения вычислительной задачи согласно закону Амдала.

Here is an explanation that a program that is programmed with concurrent patterns can be efficiently parallel.

After that the explanation of system process and system thread.

## Problems

### Deadlock

russian pdd

### Livelock

it's about trying to pass through the door

### Starvation

...

## More robust Concurrency Models

Для такой декомпозиции были придуматы несколько математических моделей (зачем?).

### Actor model (1973)

Erlang & Scala

По аналогии с философией объектно-ориентированного программирования, где каждый примитив рассматривается как объект, модель акторов выделяет в качестве универсальной сущности понятие «актора». Актор является вычислительной сущностью, которая в ответ на полученное сообщение может одновременно:

- отправить конечное число сообщений другим акторам;
- создать конечное число новых акторов;
- выбрать поведение, которое будет использоваться при обработке следующего полученного сообщения.

Не предполагается существования определённой последовательности вышеописанных действий и все они могут выполняться параллельно.

### Process Calculus Family

CCS, ACP and CSP

messages, man

### CCS (Calculus of Communicating - Robin Milner)

### CSP (Communicating sequential processes - Tony Hoare)

Formalized in 1978

![Tony Hoare pic slide]

The same dude that invented Quicksort - used currently in C, C++, Java, Python, where there is no requirement for stable sorting.

*Тут общяя теория про CSP*

**Primitives**

CSP provides two classes of primitives in its process algebra:

Events
    Events represent communications or interactions. They are assumed to be indivisible and instantaneous. They may be atomic names (e.g. on, off), compound names (e.g. valve.open, valve.close), or input/output events (e.g. mouse?xy, screen!bitmap).
Primitive processes
    Primitive processes represent fundamental behaviors: examples include STOP (the process that communicates nothing, also called deadlock), and SKIP (which represents successful termination).

[here is a slide of csp_syntax]

The syntax of CSP defines the “legal” ways in which processes and events may be combined. Let e be an event, and X be a set of events. Then the basic syntax of CSP can be defined as:

**message passing via channels**

Согласно CSP, сначала вводится множество элементарных событий (алфавит), затем из них конструируются процессы, причём из только что описанных процессов можно строить новые. Процессы, протекающие параллельно, обмениваются информацией, используя безбуферный обмен информацией типа «рандеву» между парой (и только парой) процессов посредством специального объекта — канала. При взаимодействии тот участник обмена, который обратился к каналу первым, ожидает готовности партнёра (точки рандеву); при наступлении последней инициируется обмен. Использование общей для нескольких параллельных процессов памяти в CSP не допускается.

Эта теория привела к созданию occam в 1983 британской компанией для своих же новых Транспьютеров - элемент построения многопроцессорных систем

**Comparison with the actor model**

- CSP processes are anonymous, while actors have identities.
- CSP uses explicit channels for message passing, whereas actor systems transmit messages to named destination actors. These approaches may be considered duals of each other, in the sense that processes receiving through a single channel effectively have an identity corresponding to that channel, while the name-based coupling between actors may be broken by constructing actors that behave as channels.
- CSP message-passing fundamentally involves a rendezvous between the processes involved in sending and receiving the message, i.e. the sender cannot transmit a message until the receiver is ready to accept it. In contrast, message-passing in actor systems is fundamentally asynchronous, i.e. message transmission and reception do not have to happen at the same time, and senders may transmit messages before receivers are ready to accept them. These approaches may also be considered duals of each other, in the sense that rendezvous-based systems can be used to construct buffered communications that behave as asynchronous messaging systems, while asynchronous systems can be used to construct rendezvous-style communications by using a message/acknowledgement protocol to synchronize senders and receivers.



### Languages supporting concurrent programming (в той или иной форме)

Make a slide to display them side by side.

- Ada
- Alef
- Alice
- Ateji PX
- Axum
- BMDFM
- **C++** — std::thread
- Cω
- **C#** — supports concurrent computing using lock, yield, also since version 5.0 async and - await keywords introduced
- Clojure—modern, functional dialect of Lisp on the Java platform
- Concurrent Clean
- Concurrent Collections
- Concurrent ML
- Concurrent Pascal
- Curry
- D
- E
- **ECMAScript**
- Eiffel
- Elixir
- Erlang
- FAUST
- Fortran
- **Go** — for system programming, with a concurrent programming model based on CSP
- **Haskell** — shared memory
- Hume
- Io
- Janus
- **Java** — thread class or Runnable interface
- Julia
- JoCaml
- Join Java
- Joule
- Joyce
- LabVIEW
- Limbo
- MultiLisp
- Modula-2
- Modula-3
- Newsqueak
- **occam** — influenced heavily by communicating sequential processes (CSP)
- occam-π
- Orc
- Oz-Mozart
- ParaSail
- Pict
- Raku
- **Python** — uses thread-based parallelism and process-based parallelism
- Reia
- Red/System
- **Rust**
- Scala
- SequenceL
- SR
- SuperPascal
- Unicon
- TNSDL
- VHSIC
- XC

## Now Go

A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.
Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. Their design hides many of the complexities of thread creation and management.

Each concurrent piece of code (goroutine) does not need to share memory ( shared variables with locks) with other goroutine like how you would typically approach concurrency in most other languages. Instead, it can share memory via communicating with other goroutines ( this is done by sending the data from one goroutine to another via the go channels). By default, when you pass a message, the two goroutines wait till the message is received on the other end.

The fact that Go supports this natively , covering all the memory management and thread scheduling, makes concurrency easy to implement and efficient in practice from a developer's prospective.


## Asynchronous

### But first - what is stack in synchronous code

### Then event loop

Steal from here with slides

https://iximiuz.com/en/posts/explain-event-loop-in-100-lines-of-code/

The most obvious examples are asynchronous I/O, UIs. Because we can't wait. More here on that...

Video or picture of rollercoaster

### In python specifically

https://www.youtube.com/watch?v=tSLDcRkgTsY

What is synchronous code?

Synchronous code:

Is what you're used to!
Runs functions one after another

What is asynchronous code?

Asynchronous code:

Runs multiple functions seemingly in parallel
    In a single process
    Without threads
Requires cooporative, well-behaving functions
    Functions that regularly suspend by awaiting something
Should not use blocking functions!
    No time.sleep()
    No socket.*
    Etc.
    asyncio provides non-blocking alternatives for many of these functions

A note on Python versions

The async and await keywords were introduced in Python 3.5
They are syntactic sugar on top of the asyncio module that was introduced in Python 3.4

Python 3.3 and earlier do not support this

But generator coroutines can do some of the same things

# Drawbacks
