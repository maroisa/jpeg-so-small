<script>
    export let dirPath
    export let openDirectory
    export let label = ""
</script>

<span class="font-medium text-lg">
    {label} directory:
</span>
<button class="h-12 card mb-2" on:click={() => openDirectory(label)}>
    <span class="text-blue-100/50">
        {dirPath ? dirPath : "Select " + label +" directory..."}
    </span>
</button>
