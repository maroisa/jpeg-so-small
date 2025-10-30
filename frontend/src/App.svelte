<script>
    import { OpenDirectoryDialog } from "../wailsjs/go/main/App.js";
    import DirInput from "./DirInput.svelte";

    let inputDir = null;
    let outputDir = null;

    let nameFix = "prefix"

    async function openDirectory(label) {
        let filename = await OpenDirectoryDialog();
        if (label == "input"){
            inputDir = filename
        } else if (label == "output"){
            outputDir = filename
        }
    }

    function compress(){
        window.runtime.LogPrint("From: " + inputDir + " To: " + outputDir)
    }
</script>

<main class=" bg-[#21222b] text-blue-100 p-6 flex flex-col gap-6">
    <DirInput label="input" dirPath={inputDir} openDirectory={openDirectory} />
    <DirInput label="output" dirPath={outputDir} openDirectory={openDirectory} />
    {#if inputDir != null && inputDir == outputDir}
    <div>
        <span>seems like you has same input and output directory. Add</span>
        <select class="p-2 card"on:change={(event) => nameFix = event.target.value}>
            <option value="prefix">Prefix</option>
            <option value="suffix">Suffix</option>
        </select>
        <span>for your files</span>
            <div>
                {#if nameFix == "suffix"}
                    <span class="text-blue-100/60">file1</span>
                    <input value=".min" class="card mt-4 p-2" type="text">
                    <span class="text-blue-100/60">.jpg</span>
                {:else if nameFix == "prefix"}
                    <input value="min." class="card mt-4 p-2" type="text">
                    <span class="text-blue-100/60">file1.jpg</span>
                {/if}
            </div>
        </div>
    {/if}
    <div class="grow card bg-transparent p-4 text-blue-100/50 overflow-auto">
        <p class="text-center">-- Log --</p>
        <hr class="text-blue-100/10 my-2 overflow-auto" />
        <p></p>
    </div>
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
