<script lang="ts">
    import { createSearchStore, searchHandler } from '$lib/stores/search.js';
    import { onDestroy } from 'svelte';

    export let data;
    console.log(data)

    const searchPlaces = data.places.Data.map((place: any) => ({
        ...place,
        searchTerms: `${place.Area} ${place.Postal_code} ${place.District} ${place.Settlement} ${place.Region}`
    }))

    const searchStore = createSearchStore(searchPlaces)

    const unsubscribe = searchStore.subscribe((model) => searchHandler(model))

    onDestroy(() => {
        unsubscribe();
    })
</script>

<main class="max-w-6xl mx-auto">
    <section class="my-5 text-center">
          <h1 class="text-2xl tracking-tight text-gray-900 ">Postal codes at your fingertips</h1>
    </section>
    
    <form class="max-w-md mx-auto my-10">   
        <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
        <div class="relative">
            <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
                </svg>
            </div>
            <input type="search" id="default-search" bind:value={$searchStore.search} class="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-100 dark:border-gray-100 dark:placeholder-gray-400 dark:text-gray-500 dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search Area..." required />
            <!--<button type="submit" class="text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Search</button>-->
        </div>
    </form>
    
    <div class="relative overflow-x-auto">
        <table class="rounded-md shadow-md w-50 text-sm text-left rtl:text-right text-gray-500 dark:text-gray-500 mx-auto">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-100 dark:text-gray-900">
                <tr>
                    <th scope="col" class="px-6 py-3">
                        Area
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Postal code
                    </th>
                    <th scope="col" class="px-6 py-3">
                        District
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Region
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Settlement
                    </th>
                </tr>
            </thead>
            <tbody>
                {#each $searchStore.filtered as place }
                    <tr class="odd:bg-gray border-b odd:dark:bg-gray-300 odd:dark:border-gray-300">
                        <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-gray-500">
                            {place.Area}
                        </th>
                        <td class="px-6 py-4">
                            {place.Postal_code}
                        </td>
                        <td class="px-6 py-4">
                            {place.District}
                        </td>
                        <td class="px-6 py-4">
                            {place.Region}
                        </td>
                        <td class="px-6 py-4">
                            {place.Settlement}
                        </td>
                    </tr>
                {/each}       
            </tbody>
        </table>
        <div class="flex justify-center my-6">
            <ul class="list-style-none mb-6 flex">
                <li>
                    {#if data.places.Page < 2}
                        <a
                        class="pointer-events-none relative block rounded bg-transparent px-3 py-1.5 text-sm text-surface/50 transition duration-300 dark:text-neutral-400"
                        href="#!"
                        >Previous</a
                        >
                    {:else}
                        <a
                        class="relative block rounded bg-transparent px-3 py-1.5 text-sm text-surface/50 transition duration-300 dark:text-neutral-400"
                        href="#!"
                        >Previous</a
                        >
                    {/if}
                </li>
                <li>
                  <a
                    class="relative block rounded bg-transparent px-3 py-1.5 text-sm text-gray-700 text-surface transition duration-300 hover:bg-neutral-100 "
                    href="#!"
                    >1</a
                  >
                </li>
                <li aria-current="page">
                  <a
                    class="relative block rounded bg-[#0f172a] px-3 py-1.5 text-[#6590d5] text-sm font-medium transition duration-300 focus:outline-none"
                    href="#!"
                    >2
                    <span
                      class="absolute -m-px h-px w-px overflow-hidden whitespace-nowrap border-0 p-0 [clip:rect(0,0,0,0)]"
                      >(current)</span
                    >
                  </a>
                </li>
                <li>
                  <a
                    class="relative block rounded bg-transparent px-3 py-1.5 text-sm text-gray-700 text-surface transition duration-300 hover:bg-neutral-100"
                    href="#!"
                    >3</a
                  >
                </li>
                <li>
                  <a
                    class="relative block rounded bg-transparent px-3 py-1.5 text-sm text-surface transition duration-300 hover:bg-neutral-100 focus:bg-neutral-100 focus:text-primary-700 focus:outline-none active:bg-neutral-100 active:text-primary-700 dark:text-white dark:hover:bg-neutral-700 dark:focus:bg-neutral-700 dark:focus:text-primary-500 dark:active:bg-neutral-700 dark:active:text-primary-500"
                    href="#!"
                    >Next</a
                  >
                </li>
              </ul>
        </div>
    </div>
</main>

