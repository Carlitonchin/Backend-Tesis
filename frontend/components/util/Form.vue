<script setup>
    const {inputs, workerModel} = defineProps(["inputs", "workerModel"])
    const fontSizeLabel = '1.5rem'
    const fontSizeInput = '1.3rem'
</script>

<template>
    <form id="s-form">
        <div id="s-container-container" class="flex-vertical">
        <div id="s-container-form" class="flex-horizontal flex-space-between">
        <div id="s-container-labels" class="flex-vertical flex-align-right flex-space-between">
        <label :style="{fontSize : `${fontSizeLabel}`}" class="s-label" v-for="item in inputs" :key="item.id" :for="item.id">
            {{item.label + ": "}}
        </label>
        </div>

        <div id="s-container-inputs" class="flex-vertical flex-space-between">
            <div v-for="item in inputs" :key="item.id" class="s-input">
                <UtilInput
                :type="item.type"
                :placeholder="item.placeholder"
                :model="item.model"
                :id="item.id"
                :fontSize="fontSizeInput"
                />
            </div>

        </div>

        </div>

        <div id="s-container-check" class="flex-horizontal flex-center">

        <input type="checkbox" id="worker" @change="e=>workerModel.value = e.target.value"/>
        <label id="label-worker" for="worker" :style="{fontSize : `${fontSizeInput}`}">Soy trabajador de Matcom</label>
        </div>

        </div>

    </form>
</template>

<style scoped>

    #s-container-container{
        height: 100%;
    }

    #s-form{
        height: 100%;
    }

    #s-container-form{
        height: 100%;
    }
    .s-input{
        margin-top: 0.15rem;
        width:98%;
        align-self: flex-end;
    }

    #s-container-labels{
        width: fit-content;
        min-width: fit-content;
    }

    #s-container-inputs{
        width:100%;
    }

#s-container-check{
    margin-top: 1.3rem;

}
    #worker{
        width: 1rem;
        height: 1rem;
        margin-right: 0.3rem;
        cursor: pointer;
    }

    #label-worker{
        cursor: pointer;
    }


</style>
