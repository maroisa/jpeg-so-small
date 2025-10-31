<script>
    export let active
</script>

<div hidden={!active} class="fixed top-0 left-0 size-full bg-black/20 flex justify-center items-center">
    <div class="m-20 p-4 card bg-base-100 text-blue-100/50 flex flex-col max-w-xl w-full transition-all duration-100">
        <p class="text-sky-100 text-center">Success!</p>
        <button on:click={() => active = false} class="
            w-full
            card
            p-2
            mt-4
            bg-sky-900/10
            hover:bg-sky-900/20
            active:bg-sky-900/30
            text-sky-400/60
            hover:text-sky-400/80
            active:text-sky-400
            border-sky-400/60
            hover:border-sky-400/80
            active:border-sky-400"
        >
            Ok
        </button>
    </div>
</div>