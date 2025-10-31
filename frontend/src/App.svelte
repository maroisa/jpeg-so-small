<script>
    import { OpenDirectoryDialog, Compress } from "../wailsjs/go/main/App.js";
    import DirInput from "./DirInput.svelte";
    import ProgressPopup from "./ProgressPopup.svelte";
    import SuccessPopup from "./SuccessPopup.svelte";

    let qualityValue = 90;
    let inputDir = null;
    let outputDir = null;

    let nameFixType = "prefix"
    $: nameFix = nameFixType == "prefix" ? "min." : ".min"

    let progressValue = ""
    let successActive = false

    window.runtime.EventsOn("onCompressed", (inputPath, outputPath) => {
        if (inputPath == "" && outputPath == ""){
            progressValue = ""
            return
        }
        progressValue = inputPath + " > " + outputPath
    })

    async function openDirectory(label) {
        let filename = await OpenDirectoryDialog();
        if (label == "input"){
            inputDir = filename
        } else if (label == "output"){
            outputDir = filename
        }
    }

    async function compress(){
        if (inputDir != outputDir){
            nameFix = ""
            nameFixType = ""
        }
        await Compress(inputDir, outputDir, qualityValue, nameFixType, nameFix)
        
        successActive = true
    }
</script>

<SuccessPopup active={successActive} />
<ProgressPopup value={progressValue} />

<main class=" bg-[#21222b] text-blue-100 p-6 flex flex-col gap-4">
    <DirInput label="input" dirPath={inputDir} openDirectory={openDirectory} />
    <DirInput label="output" dirPath={outputDir} openDirectory={openDirectory} />
    <p class="text-lg font-medium">Quality: {qualityValue}</p>
    <input bind:value={qualityValue} type="range" min="1" max="100" class="w-full">
    {#if inputDir != null && inputDir == outputDir}
        <div>
            <span>seems like you has same input and output directory. Add</span>
            <select class="p-2 card"on:change={(event) => nameFixType = event.target.value}>
                <option value="prefix">Prefix</option>
                <option value="suffix">Suffix</option>
            </select>
            <span>for your files</span>
            <div>
                {#if nameFixType == "suffix"}
                    <span class="text-blue-100/60">file1</span>
                    <input bind:value={nameFix} class="card mt-4 p-2" type="text">
                    <span class="text-blue-100/60">.jpg</span>
                {:else if nameFixType == "prefix"}
                    <input bind:value={nameFix} class="card mt-4 p-2" type="text">
                    <span class="text-blue-100/60">file1.jpg</span>
                {/if}
            </div>
        </div>
    {/if}

    <div class="grow"></div>
    <button 
        on:click={compress}
        disabled={!inputDir || !outputDir}
        class="
            h-12 card
            disabled:bg-cyan-100/10
            disabled:text-white/10
            bg-cyan-100/60
            hover:bg-cyan-100/70
            active:bg-cyan-100
            text-slate-800
            font-bold text-lg">
        Compress
    </button>
</main>
