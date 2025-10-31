<script>
    export let value
</script>

<div hidden={!value} class={"absolute top-0 left-0 size-full bg-black/10 flex flex-col justify-center items-center"}>
    <div class=" m-20 p-4 card bg-base-100 text-blue-100/50 flex flex-col max-w-xl w-full">
        <p>{value}</p>
        <button class="
            w-full
            card
            p-2
            mt-4
            bg-red-900/10
            hover:bg-red-900/20
            active:bg-red-900/30
            text-red-400/60
            hover:text-red-400/80
            active:text-red-400
            border-red-400/60
            hover:border-red-400/80
            active:border-red-400"
        >
            Cancel
        </button>
    </div>

</div>