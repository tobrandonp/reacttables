# Svelte Grokking


## Reactive Declarations
#### Simple:
        $: doubled = count * 2
        $: means this code will be run any time count changes

##### You can even have checks using reactivity:

        let count = 0;
        $: if (count >= 5){
            alert("OH NO")
        }
            
        function handleClick() {
            count += 1;
        }

-------------------------------------------------
#### Summing Numbers using Reactivity:
        let numbers = [1, 2, 3, 4];

        function addNumber() {
            numbers = [...numbers, numbers.length + 1];    // Way #1 - Idiomatic
            numbers[numbers.length] = numbers.length + 1;  // Way #2, oddly enough*****
            numbers = numbers;                             // Also works, but not idiomatic
        }

        $: sum = numbers.reduce((total, currentNumber) => total + currentNumber, 0);

        <p>{numbers.join(' + ')} = {sum}</p>

        <button on:click={addNumber}>
            Add a number
        </button>

***** Assignments to properties of arrays and objects — e.g. obj.foo += 1 or array[i] = x — work the same way as assignments to the values themselves.



## Components
        <script>
            import MyComponent from './MyComponent.svelte'
        <script>

        <div>
            <MyComponent />
        </div>



## Props (a la - passing variables down to other components)
    Nested.svelte
        <script>
            export let answer = "unknown!";  // default prop value
        </script>

        <p>The answer is {answer}</p>

    App.svelte:
        <script>
            import Nested from './Nested.svelte';
        </script>

        <Nested answer={42} />
        <Nested/> <!-- Prints default value -->


## Spread Props
### If the prop values match the attribute values (E.g. <div a={a} b={b}>), we can use spread props
        PackageInfo.svelte:
        <script>
            import PackageInfo from './PackageInfo.svelte';

            const pkg = {
                name: 'svelte',
                speed: 'blazing',
                version: 4,
                website: 'https://svelte.dev'
            };
        </script>

        <PackageInfo
            {...pkg}  <!-- Spread operator -->
        />
        <!-- Same as: -->
        <PackageInfo
            name={pkg.name}
            speed={pkg.speed}
            version={pkg.version}
            website={pkg.website}
        />


    Conversely, if you need to reference all the props that were passed into a component, including ones that weren't declared with export, you can do so by accessing $$props directly. It's not generally recommended, as it's difficult for Svelte to optimise, but it's useful in rare cases.


## Conditional Rendering of HTML elements

        <script>
            let count = 0;

            function increment() {
                count += 1;
            }
        </script>

        <button on:click={increment}>
            Clicked {count}
            {count === 1 ? 'time' : 'times'}
        </button>

        {#if count > 10}
            <p>{count} is greater than 10</p>
        {:else if count < 5}
            <p>{count} is less than 5</p>
        {:else}
            <p>{count} is between 0 and 10</p>
        {/if}


#### A # character always indicates a block opening tag. A / character always indicates a block closing tag. A : character, as in {:else}, indicates a block continuation tag.


## Looping to generate Elements
    Instead of manually writing a button for each color, we can use {#each} and {/each}
        <script>
            const colors = ['red', 'orange', 'yellow', 'green', 'blue', 'indigo', 'violet'];
            let selected = colors[0];
        </script>

        <h1 style="color: {selected}">Pick a colour</h1>

        <div>
            {#each colors as color, i}
                <button
                    aria-current={selected === color}
                    aria-label=color
                    style="background: {color}"
                    on:click={() => selected = color}
                >{i+1}</button>
            {/each}

        </div>

        <style>
            h1 {
                transition: color 0.2s;
            }

            div {
                display: grid;
                grid-template-columns: repeat(7, 1fr);
                grid-gap: 5px;
                max-width: 400px;
            }

            button {
                aspect-ratio: 1;
                border-radius: 50%;
                background: var(--color, #fff);
                transform: translate(-2px,-2px);
                filter: drop-shadow(2px 2px 3px rgba(0,0,0,0.2));
                transition: all 0.1s;
            }

            button[aria-current="true"] {
                transform: none;
                filter: none;
                box-shadow: inset 3px 3px 4px rgba(0,0,0,0.2);
            }
        </style>

## Event Modifiers
        <button on:click|once={() => alert('clicked')}>
            Click me
        </button>

    -preventDefault — calls event.preventDefault() before running the handler. Useful for client-side form handling, for example.
    -stopPropagation — calls event.stopPropagation(), preventing the event reaching the next element
    -passive — improves scrolling performance on touch/wheel events (Svelte will add it automatically where it's safe to do so)
    -nonpassive — explicitly set passive: false
    capture — fires the handler during the capture phase instead of the bubbling phase
    -once — remove the handler after the first time it runs
    -self — only trigger handler if event.target is the element itself
    -trusted — only trigger handler if event.isTrusted is true, meaning the event was triggered by a   user action rather than because some JavaScript called element.dispatchEvent(...)


## Component Events

##### App.svelte:
        <script>
            import Inner from './Inner.svelte';

            function handleMessage(event) {
                alert(event.detail.text);
            }
        </script>

        <Inner on:message={handleMessage} />

##### Inner.svelte


        <script>
            import { createEventDispatcher } from 'svelte';

            const dispatch = createEventDispatcher();
            
            function sayHello() {
                dispatch('message', {
                    text: 'Hello!'
                });
            }
        </script>

        <button on:click={sayHello}>
            Click to say hello
        </button>


### This is useful because:
#### 1. Form Components:

Imagine you're building a complex form with multiple reusable input components like custom dropdowns, date pickers, or sliders. Each of these components might manage its internal state but needs to communicate changes up to the parent form component.

    Child Component (e.g., CustomDropdown.svelte): Dispatches an event when a user selects an option.
    Parent Component (e.g., Form.svelte): Listens to the event and updates the form's overall state or triggers validation.

#### 2. Dialogs or Modals:

You might have a generic modal component that can be used throughout your application. The modal might have actions like 'Save', 'Cancel', or 'Delete'.

    Child Component (e.g., Modal.svelte): Dispatches events for each action, like on:save, on:cancel, or on:delete.
    Parent Component: Listens for these events and handles them appropriately, maybe saving data, closing the modal, or asking for extra confirmation when deleting.

#### 3. List Item Actions:

If you have a list where each item has actions (like 'edit', 'delete', 'view details'), you can use event dispatching to handle these actions in the parent list component.

    Child Component (e.g., ListItem.svelte): Dispatches an event when an action is taken on an item.
    Parent Component (e.g., List.svelte): Listens for these events and handles them, perhaps by showing a detailed view, editing the item, or removing it from the list.

#### 4. Notifications or Messages:

If components in your application need to notify the user about certain events (like a successful action, a warning, or an error), you can use a notification system based on event dispatching.

    Child Component: Dispatches an event with the notification details.
    Parent Component or Notification Manager: Listens for these events and displays notifications to the user.

#### 5. Interaction Between Disparate Components:

Sometimes, you might have components that are not directly related or nested but need to communicate. You can use an event dispatching mechanism to bubble up events to a common ancestor, which then delegates the action to another component.

    Child Component A: Dispatches an event.
    Common Ancestor Component: Listens for the event and triggers a method in Child Component B based on the event details.






## Event Forwarding

Unlike DOM events, component events don't bubble. If you want to listen to an event on some deeply nested component, the intermediate components must forward the event.


### Example

App.svelte:

        <script>
            import Outer from './Outer.svelte';

            function handleMessage(event) {
                alert(event.detail.text);
            }
        </script>

        <Outer on:message={handleMessage} />


Inner.svelte:

<script>
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function sayHello() {
		dispatch('message', {
			text: 'Hello!'
		});
	}
</script>

    <script>
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function sayHello() {
		dispatch('message', {
			text: 'Hello!'
		});
	}
    </script>

    <button on:click={sayHello}>
        Click to say hello
    </button>


Outer.svelte

    <script>
        import Inner from './Inner.svelte';
    </script>

    <Inner on:message />  => This line forwards the Inner message



#### Example 2:

##### In order for us to be able to pass the onClick event to the button, we need to include the on:click inside BigRedButton.svelte Button element

App.svelte:

    <script>
        import BigRedButton from './BigRedButton.svelte';
        import horn from './horn.mp3';

        const audio = new Audio();
        audio.src = horn;

        function handleClick() {
            audio.load();
            audio.play();
            alert("PLAYED HORN")
        }
    </script>

    <BigRedButton on:click={handleClick} />



BigRedButton.svelte:

    <button on:click>
        Push
    </button>

    <style>
        button {
            font-size: 1.4em;
            width: 6em;
            height: 6em;
            border-radius: 50%;
            background: radial-gradient(circle at 25% 25%, hsl(0, 100%, 50%) 0, hsl(0, 100%, 40%) 100%);
            box-shadow: 0 8px 0 hsl(0, 100%, 30%), 2px 12px 10px rgba(0,0,0,.35);
            color: hsl(0, 100%, 30%);
            text-shadow: -1px -1px 2px rgba(0,0,0,0.3), 1px 1px 2px rgba(255,255,255,0.4);
            text-transform: uppercase;
            letter-spacing: 0.05em;
            transform: translate(0, -8px);
            transition: all 0.2s;
        }

        button:active {
            transform: translate(0, -2px);
            box-shadow: 0 2px 0 hsl(0, 100%, 30%), 2px 6px 10px rgba(0,0,0,.35);
        }
    </style>


horn.mp3:

    <base64 string...>





## Binding A Value - Text Inputs

    <script>
	let name = 'world';
    </script>

    <input bind:value={name} />

    <h1>Hello {name}!</h1>


## Checkbox:
    <script>
	let yes = false;
    </script>

    <label>
        <input type="checkbox" bind:checked={yes} />
        Yes! Send me regular email spam
    </label>

    {#if yes}
        <p>
            Thank you. We will bombard your inbox and sell
            your personal details.
        </p>
    {:else}
        <p>
            You must opt in to continue. If you're not
            paying, you're the product.
        </p>
    {/if}

    <button disabled={!yes}>Subscribe</button>



## Binding With a slider:



    <script>
        let a = 1;
        let b = 2;
    </script>

    <label>
        <input type="number" bind:value={a} min="0" max="10" />
        <input type="range" bind:value={a} min="0" max="10" />
    </label>

    <label>
        <input type="number" bind:value={b} min="0" max="10" />
        <input type="range" bind:value={b} min="0" max="10" />
    </label>

    <p>{a} + {b} = {a + b}</p>


## SUBMIT EXAMPLE & Bind Selection

    <script>
        let questions = [
            {
                id: 1,
                text: `Where did you go to school?`
            },
            {
                id: 2,
                text: `What is your mother's name?`
            },
            {
                id: 3,
                text: `What is another personal fact that an attacker could easily find with Google?`
            }
        ];

        let selected;

        let answer = '';

        function handleSubmit() {
            alert(
                `answered question ${selected.id} (${selected.text}) with "${answer}"`
            );
        }
    </script>

    <h2>Insecurity questions</h2>

    <form on:submit|preventDefault={handleSubmit}>
        <select
            bind:value={selected}
            on:change={() => (answer = '')}
        >
            {#each questions as question}
                <option value={question}>
                    {question.text}
                </option>
            {/each}
        </select>

        <input bind:value={answer} />

        <button disabled={!answer} type="submit">
            Submit
        </button>
    </form>

    <p>
        selected question {selected
            ? selected.id
            : '[waiting...]'}
    </p>

## Bind Groups
If you have multiple type="radio" or type="checkbox" inputs relating to the same value, you can use bind:group along with the value attribute. Radio inputs in the same group are mutually exclusive; checkbox inputs in the same group form an array of selected values.



# Type Coersion 
Svelte handles type conversion using type="number" attributes

