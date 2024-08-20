    <script>
    import toast, { Toaster } from 'svelte-french-toast';
    import { fade } from 'svelte/transition';
    import { fly } from 'svelte/transition';

    let username = '';
    let password = '';
    let confirmPassword = '';

    function passwordsMatch() {
        if (password !== confirmPassword) {
            toast.error("Passwords do not match!");
            return false;
        } else {
            return true;
        }
    }

    function name(params) {
      
    }

    async function submit(event) {
        event.preventDefault();

        if (passwordsMatch()) {
            const fetchPromise = fetch("http://127.0.0.1:3000/api/signup", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    username: username,
                    password: password
                })
            }).then(async (res) => {
                if (res.ok) {
                    const data = await res.json();
                    return data;
                } else {
                    const errorData = await res.json();
                    throw new Error(errorData.error || 'Unknown error occurred');
                }
            });

            toast.promise(
                fetchPromise,
                {
                    loading: 'Creating account...',
                    success: 'Account created successfully!',
                    error: (err) => `Unable to connect to server`
                }
            ).catch((error) => {
                console.error("Unexpected error:", error);
            });
        }
    }
</script>

    <!-- Background -->
    <div>
        <div class="absolute bottom-auto left-auto -right-14 -top-2 h-[500px] w-[500px] -translate-x-[30%] translate-y-[20%] rounded-full bg-[rgba(173,109,244,0.5)] opacity-50 blur-[80px]"></div>
        <div class="absolute flex left-96 h-[500px] w-[500px] -translate-x-[30%] translate-y-[20%] rounded-full bg-[rgba(173,109,244,0.5)] opacity-50 blur-[80px]"></div>
        <div class="absolute bottom-auto left-auto left-96 -top-2 h-[500px] w-[500px] -translate-x-[30%] translate-y-[20%] rounded-full bg-[rgba(173,109,244,0.5)] opacity-50 blur-[80px]"></div>
        <div class="absolute bottom-auto left-auto right-96 -top-2 h-[500px] w-[500px] -translate-x-[30%] translate-y-[20%] rounded-full bg-[rgba(173,109,244,0.5)] opacity-50 blur-[80px]"></div>

    </div>
    <div class="relative z-10 flex flex-col items-center px-6 mx-auto mt-16 sm:h-auto lg:py-0">
        <div in:fade={{ duration: 400 }} class="w-full bg-white rounded-lg shadow-xl shadow-[#e7d5fb] sm:max-w-md xl:p-0 border border-gray-500">
            <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl">
                    Create an account
                </h1>
                <form class="space-y-4 md:space-y-6" action="#">
                    <div>
                        <label for="username" class="block mb-2 text-sm font-medium text-gray-900">Your Username</label>
                        <input
                            type="username"
                            name="username"
                            id="username"
                            bind:value={username}
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 hover:border-primary-600 block w-full p-2.5"
                            placeholder="coolguy512"
                            required
                        />
                    </div>
                    <div>
                        <label for="password" class="block mb-2 text-sm font-medium text-gray-900">Password</label>
                        <input
                            type="password"
                            name="password"
                            id="password"
                            bind:value={password}
                            placeholder="••••••••"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                            required
                        />
                    </div>
                    <div>
                        <label for="confirm-password" class="block mb-2 text-sm font-medium text-gray-900">Confirm password</label>
                        <input
                            type="password"
                            name="confirm-password"
                            id="confirm-password"
                            bind:value={confirmPassword}
                            placeholder="••••••••"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                            required
                        />
                    </div>
                    <button on:click={submit}
                        type="submit"
                        class="transition hover:bg-gradient-to-r w-full text-black border border-[#AD6DF4] hover:from-[#AD6DF4] hover:to-purple-600 hover:text-white focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
                    >
                        Submit
                    </button>
                    <p class="text-sm font-light text-gray-500">
                        Already have an account? <a href="/login" class="font-medium text-primary-600 hover:underline">Login</a>
                    </p>
                </form>
            </div>
        </div>
</div>
