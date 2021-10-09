TODO put it on your own server in case the projector does not connect

TODO remove git markers, bracket colorizer (write down changed settings)

TODO set PS1 to minimal

# Concurrency Patterns

## Introduction (up to pigeons)

// first slide with crosswalk

At first some background. So if you look around in the world at large, what you see is a lot of independently executing things. You see people in the audience doing their own thing, there's people outside, there's cars going by. All of those things are independent agents, if you will.

// second slide with people in some arena

И все эти агенты общаются между собой и действуют в единой среде.

// slide with elevators and with pigeons

Причем состояние в котором окажутся эти агенты может быть неопределена или зависеть от множества факторов вне нашего контроля.

// slide with one-way road sign

И если вы думаете о написании компьютерной программы, если вы хотите моделировать эту среду или взаимодействовать с ней - единое последовательное выполнение - не очень хороший подход.

Поэтому люди стали придумывать способы моделировать такие системы, но прежде нужно сделать важное замечание,

// slide concurrency is not parallelism

### Concurrency is not parallelism

And so cuncurrency is really a way of writing or structuring you program to deal with real world. And what I mean by concurrency is the composition of independently executing computations (usually functions, but they don't have to be). Concurrency is a way to structure software. It is not parallelism.

// slide with cpu and threads

Выполнение может быть не параллельным, но при этом конкурентным, как например на системе с одним процессором - в один момент времени процессор выплняет одну инструкцию, но при этом эти инструкции чередуются.

// slide with airport 1 and 2

### Airport example

Хорошим примером может служить аэропорт с одной взлетно посадочной полосой. Где наша сисема конкурентна, т.к. представлена множеством самолётов. Но при этом не параллельна - только один самолет может использовать полосу. Но такая канкурентная система при этом очень эффективна - в большинстве аэропортов используется только одна полоса. Но при этом остальные самолеты не стоят - они готовятся к тому чтобы использовать эту полосу.

// slide with parallel plane landing

С другой стороны, хорошо написанная параллельная программа может эффективно работать параллельно в многопроцессорной или распределенной среде.

Это свойство важно и источник заблуждений, потому что исторически люди сначала задумывались как бы сделать свою линейную программу быстрее с помощью декомпозиции на части (т.е. сделать программу конкурентной) и запустить на нескольких процессорах. Поэтому часто конкурентность (метод декомпозиции и распределения) отожествлялся с параллелизмом (т.е. непосредственным выполнением нескольких вещей в один момент времени). Только пожже люди начали задумываться над более широким применением конкурентности как методом моделирования программ. Но сначало, зачем вообще нужно распараллеливание вычислений - почему мы не можем просто увеличивать мощность.

// slide with moore'd law

### But why do we even need parallel?

В последнее время, чтобы получить возможность задействовать на практике ту дополнительную вычислительную мощность, которую предсказывает закон Мура, стало необходимо задействовать параллельные вычисления. На протяжении многих лет производители процессоров постоянно увеличивали тактовую частоту и параллелизм на уровне инструкций, так что на новых процессорах старые однопоточные приложения исполнялись быстрее без каких-либо изменений в программном коде.[8] Примерно с середины десятилетия 2000-х годов по разным причинам производители процессоров предпочитают многоядерные архитектуры, и для получения всей выгоды от возросшей производительности ЦП программы должны переписываться в соответствующей манере. Однако не каждый алгоритм поддается распараллеливанию, определяя, таким образом, фундаментальный предел эффективности решения вычислительной задачи согласно закону Амдала.

Шкала у этого графика логарифмическая, а так это вообще экспонента

// slide with exponential

// slide with imb supercomputers

Ок, первая причина конкурентного моделирования - это чтобы потом распараллелить нашу программу. Хочу уже написать что-нибудь. Но,

// slide with all languages

### Languages supporting concurrent programming

Все эти языки в той-или иной форме имеют способы конкурентного разделения задач.
Но конкурентность можно смоделировать на разных уровнях и благо для нашей цели подойдет системный уровень - во многих системах уже реализованы методы разделения.

// slide with system-level classification of concurrency

### Concurrency at system level

At the operating system level:

- Threads
- Processes

// slide with threading is better than multiprocessing

Для программ с узким назначением многопоточность лучше чем множество процессов.

// slide with items per process and per thread

Обяснение этому простое - процесс это более высокая абстракция на потоком.

**Процесс**

На уровне ядра процесс содержит один или несколько потоков, которые совместно используют ресурсы процесса, такие как память и дескрипторы файлов: процесс - это единица ресурсов, а поток - это единица планирования и выполнения. Планирование ядра обычно выполняется заранее или, что реже, совместно. Процессы обычно являются многозадачными с вытеснением, а переключение процессов является относительно дорогостоящим, помимо базовой стоимости переключения контекста.

Пример разделения процессов на about:processes в firefox.

**Поток (thread)**

Поток - это «облегченная» единица планирования ядра. В каждом процессе существует по крайней мере один поток. Если в процессе существует несколько потоков, они используют одну и ту же память и файловые ресурсы. Потоки не владеют ресурсами, кроме стека и копии регистров. Процессор (cpu) (не процесс, а процессор) работает с потоками. Он выполняет в одни момент времени один поток. Переключение между потоками - это переключение контекста.

## CPython and Ruby

Теперь нам нужно кое-что убрать. Python имеет многопроцессорность для задач, которые могут выиграть от параллельного выполнения. Это потому, что у Python есть GIL, который ограничевает, что может выполняться только один поток в один момент веремени. А у Ruby есть GVL, который аналогично ограничивает то, что к виртуальной машине Ruby может быть доступ на выполнение только у одного потока. Причина этому, безопасный для потоков интерпретатор или виртуальная машина - это довольно тяжелая вещь. JVM - это потокобезопасная вещь, но она процвела через боль и страдания сначало под поддержкой Sun, пожже Oracle. (Kotlin, Scala). Ни python, ни Ruby не имели таких ресурсов.

// slide with threading

## Threading

Итак, мы поняли, что нам нужно - многопоточность - разделим. Тут показываю код. Всё классно. Но прежде, можно представить.

// slide with question

..Что если я хочу много потоков? Тогда у меня проблема ресурсов. Потоки довольно легкие, но на создание потока всё же аллоцируется около 500кб и для переключения между потоками процессор должен выполнить контекстный переход. То есть уже для 1000 потоков программе потребуется около полгигабайта и процессор не сможет достаточно быстро циклировать между потоками. То есть когда мы создаем потоки мы на самом деле ограничены довольно жесткими рамками.

// slide with ideal performance of system threads

// slide with classification

И люди создали корутины (в широком смысле слова). Корутины сделаны на уровне языка программирования гораздо легче чем системные потоки, на которые аллоцируется около 500кб при создании (наиболее экстремальным примером была java, у которого jvm до некоторого времени агрессивно аллоцировал до 1мб на создание потока).

Пример чистой реализации корутин можно показать на примере go. Тут открыть VSCode и создать базовую горутину. У горутин свой стак, который увеличивается или уменьшается по необходимости. Это не систеный поток. Они очень дешевые. На самом деле они мультиплексятся на потоки в отношении M:N.

Тут показать в VSCode, что можно легко создать 1_000_000 корутин.

Но это всё круто, но возвращаясь к вопросу - **Зачем мне нужно столько корутин**. Я их никогда в жизни не распараллелю. Но на самом деле моя цель совсем в другом, я хочу смоделировать поведение реальных агентов, о которых я говорил в начале. Я хочу сделать мою программу фундаментально конкурентной. Можно продемонстрировать на самом тупом примере.

Тут показать в VSCode,


## Sharing and communicating

Но модель не полна, потому что между агентами нет никакого общения.

// slide with Key and Peele

Для выполнения большинства задач нужно общение или синхронизация между агентами.

// slide with birds

Первое, что люди придумали - это разделение памяти. То есть использовать какой-то общий ресурс, в который можно записывать, или читать из него.

**Code it**

Но это ведет к очнь сложным проблемам.

// slide with problems

## Problems

// slide with race condition

### Race condition

Состояние гонки - возникает, когда два агента пытаются в одно и то же время использовать один общий ресурс при этом не согласовывая своих действий друг с другом.

// slide with single train

Например у нас есть поезд, при этом один путь. Если мы добавим второй поезд - получится состояние гонки.

Трудно отлаживать вручную. Пример стандартной библиотеки golang и реализованного детектора гонок.

А вообще аварии на нерегулиремых перекрестках - тоже состояние гонки. Используется общий ресурс несколькими агентами. Еще один пример:

Допустим, у вас есть аквариум с рыбками. Вы кормите рыбок каждый день после школы. Но вы не знаете, что ваша сестра делает то же самое! Когда она приходит домой, она тоже кормит рыб. Пару месяцев и рыбы погибли от перекорма. Здесь ваши незапланированные процедуры кормления были в состоянии гонки: вы не знали, кормили рыбу или нет, поэтому вы оба сделали это, потому что знали, что это должно быть сделано.

Let say that you own a tank of fishes. You feed the fishes every day after you come home from school. But the thing you don't know, is that your sister does the same! When she comes home, she feeds the fishes too. Couple of months and fishes died for overfeeding. Here your unplanned feeding routines were in race condition: you didn't know was the fishes fed or not, so you both did it because you knew that it had to be done.

// code it take the previous one

Можем взять те же горутины и представить им общий ресурс (число), к которому они будут прибавлять единицу.

// slide with semaphore

Один из способов её решить - семафор. По сути им может выступать любой индикатор, что ресурс занят. Идею можно немного усложнить и сразу добавить концепцию принадлежности. Т.е. тот, кто заблокировал индикатор, должен его и разблокировать. Пример - горорящая подушка

// slide with breaking bad

// solve race example with mutexes

Но семафоры или мьютексы могут привести к следующей проблеме.

### Deadlock

// slide with pdd

Каждый говорит, что мьютекс заблокирован машиной справа от него (или если быть совсем точным, то в данной ситуации 4 мьютекса). В итоге все стоят, хотя ресурс перекрестка свободен.

// code it

дороги - мьютексы, перекресток - слайс, машины - горутины (блокируют дорогу -> думают -> блокируют дорогу помехи). Определить функцию помехи справа.

Что такое "действовать по взаимной договоренности". Но даже если мы что-то придумаем, то ситуация становится сложнее в таком случае.

// slide with massive deadlock

И все решения, которые мы придумаем имеют риск привести к еще двум проблемам.

### Starvation

// slide with starvation

### Livelock

// slide with livelock

Это когда вы с кем-то хотите пройти через дверь и уступаете друг другу. Потом одновременно начинаете движение. Потом останавливаетесь. И в итоге вы вдвоем выглядите как придурки, потому что не можете использовать ресурс двери.


TODO you are here
## People though of something better

**here vids of people breaking computers.**

So people don't like to micromanage problems, so later they though up of ways to communicate without sharing memory, but first, problems:


## Actor model and Process calculus

In computer science, the process calculi (or process algebras) are a diverse family of related approaches for formally modelling concurrent systems. Process calculi provide a tool for the high-level description of interactions, communications, and synchronizations between a collection of independent agents or processes. They also provide algebraic laws that allow process descriptions to be manipulated and analyzed, and permit formal reasoning about equivalences between processes.

Two main approaches: **Actor model** and **Process calculus**.

There are many similarities between the two approaches, but the main difference is

- Processes in the process calculi are anonymous, and communicate by sending messages either through named channels (synchronous or asynchronous), or via ambients (which can also be used to model channel-like communications (Cardelli and Gordon 1998)). In contrast, actors in the Actor model possess an identity, and communicate by sending messages to the mailing addresses of other actors (this style of communication can also be used to model channel-like communications—see below).

### Actor model (1973)

https://rocketeer.be/articles/concurrency-in-erlang-scala/

The Actor model was inspired by the laws of physics and depends on them for its fundamental axioms, i.e. physical laws (see Actor model theory); the process calculi were originally inspired by algebra (Milner 1993).

Erlang & Scala

По аналогии с философией объектно-ориентированного программирования, где каждый примитив рассматривается как объект, модель акторов выделяет в качестве универсальной сущности понятие «актора». Актор является вычислительной сущностью, которая в ответ на полученное сообщение может одновременно:

- отправить конечное число сообщений другим акторам;
- создать конечное число новых акторов;
- выбрать поведение, которое будет использоваться при обработке следующего полученного сообщения.

Не предполагается существования определённой последовательности вышеописанных действий и все они могут выполняться параллельно.

Of importance in this model is that all communications are performed asynchronously. This implies that the sender does not wait for a message to be received upon sending it, it immediately continues its execution. There are no guarantees in which order messages will be received by the recipient, but they will eventually be delivered.

A second important property is that all communications happen by means of messages: there is no shared state between actors. If an actor wishes to obtain information about the internal state of another actor, it will have to use messages to request this information. This allows actors to control access to their state, avoiding problems like the lost-update problem. Manipulation of the internal state also happens through messages.

### Erlang

A simple application that uses an actor can be seen below. In this application, an actor is defined which acts as a basic counter. We send 100.000 increment messages to the actor and then request it to print its internal value.

Lines 1 & 2 defines the module and the exported functions. Lines 4 till 7 contain the run function, which starts a counter process and starts sending increment messages. Sending these messages happens in lines 15 till 18, using the message-passing operator (!). As Erlang is a purely functional language, it has no loop structures. Therefore, this has to be expressed using recursion. These extremely deep recursion stacks would lead to stack overflows in Java, yet Erlang is optimized for these usage patterns. The increment message in this example also carries a parameter, to show Erlangs parameter capabilities. The state of the counter is also maintained using recursion: upon receiving an inc message, the counter calls itself with the new value which causes it to receive the next message. If no messages are available yet, the counter will block and wait for the next message.

**Actor scheduling in Erlang**

Erlang uses a preemptive scheduler for the scheduling of processes [7]. When they have executed for a too long period of time (usually measured in the amount of methods invoked or the amount of CPU-cycles used), or when they enter a receive statement with no messages available, the process is halted and placed on a scheduling queue.

This allows for a large number of processes to run, with a certain amount of fairness. Long running computations will not cause other processes to become unresponsive.

### Scala

**Thread-based vs. Event-based actors**

Scala makes the distinction between thread-based and event-based actors.

Thread-based actors are actors which each run in their own JVM thread. They are scheduled by the Java thread scheduler, which uses a preemptive priority-based scheduler. When the actor enters a receive block, the thread is blocked until messages arrive. Thread-based actors make it possible to do long-running computations, or blocking I/O operations inside actors, without hindering the execution of other actors.

There is an important drawback to this method: each thread can be considered as being heavy-weight and uses a certain amount of memory and imposes some scheduling overhead. When large amounts of actors are started, the virtual machine might run out of memory or it might perform suboptimal due to large scheduling overhead.

In situations where this is unacceptable, event-based actors can be used. These actor are not implemented by means of one thread per actor, yet instead they run on the same thread. An actor that waits for a message to be received is not represented by a blocked thread, but by a closure. This closure captures the state of the actor, such that it’s computation can be continued upon receiving a message [10]. The execution of this closure happens on the thread of the sender.

Event-based actors provide a more light-weight alternative, allowing for very large numbers of concurrently running actors. They should however not be used for parallelism: since all actors execute on the same thread, there is no scheduling fairness.


**Safety in Scala concurrency**

Another potential pit-fall in Scala comes from the fact that it mixes actors with object-oriented programming. It is possible to expose the internal state of an actor through publicly available methods for retrieving and modifying this state. When doing so, it is possible to modify an object by directly invoking its methods, that is: without using messages. Doing so means that you no longer enjoy the safety provided by the actor model.

...


## Process Calculus Family

CCS, ACP and **CSP**

channels, man

### CCS (Calculus of Communicating - Robin Milner)

...

### CSP (Communicating sequential processes - Tony Hoare)

Formalized in 1978

// Tony Hoare pic slide

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

## Now Go

**Introduce challels here**

// VScode here

Go slightly extends on the idea and adds buffered channels and there is no restriction on the number of clients of the channel.

// VScode here

// main idea slide

Каждому горутине фрагменту кода не нужно делить память (общие переменные с блокировками) с другой горутиной, как вы обычно подходите к конкурентности в большинстве других языков. Вместо этого он может совместно использовать память посредством связи с другими горутинами (это делается путем отправки данных из одной горутины в другую по каналам go). По умолчанию, когда вы передаете сообщение, две горутины ждут, пока сообщение не будет получено на другом конце.

**Это основная идея каналов.**

https://go.dev/blog/codelab-share

Это может быть продемонстрированно с помощью прошлого примера с гонкой

// VScode here

**Then take the previous code with race condition and rewrite using channels**

It is interesting that with this demands the channel serves two functions - delivering message and synchronizing routines.


### Some advanced shit with channels

...

### Real life examples

...

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

### In javascript

...

http://latentflip.com/loupe/?code=JC5vbignYnV0dG9uJywgJ2NsaWNrJywgZnVuY3Rpb24gb25DbGljaygpIHsKICAgIHNldFRpbWVvdXQoZnVuY3Rpb24gdGltZXIoKSB7CiAgICAgICAgY29uc29sZS5sb2coJ1lvdSBjbGlja2VkIHRoZSBidXR0b24hJyk7ICAgIAogICAgfSwgMjAwMCk7Cn0pOwoKY29uc29sZS5sb2coIkhpISIpOwoKbHNldFRpbWVvdXQoZnVuY3Rpb24gdGltZW91dCgpIHsKICAgIGNvbnNvbGUubG9nKCJDbGljayB0aGUgYnV0dG9uISIpOwp9LCA1MDAwKTsKCmNvbnNvbGUubG9nKCJXZWxjb21lIHRvIGxvdXBlLiIpOw%3D%3D!!!PGJ1dHRvbj5DbGljayBtZSE8L2J1dHRvbj4%3D

### Drawbacks

Важно заметить, что асинхронная модель в том виде в котором она представлена в js и python - это более узкая модель по сравнению с болле общими корутинами (go и scala). Потому что всегда должен быть event loop. Т.е. тот кто всё запускает, и все функции должны добровольно участвовать в нем. С помощью такой модели невозможно в чистом виде сделать пример с перекрестком. Но напротив с помощью корутинной модели можно легко сделать модель асинхронного контроля. (TODO example in go)
