<script>
    import { onMount } from 'svelte';
    import toast, { Toaster } from 'svelte-french-toast';

    let username = '';
    let password = '';
    let confirmPassword = '';

    function passwordsMatch(){
      if(password != confirmPassword){
        toast.error("Passwords do not match!");
        return false;
      }else{
        return true;
      }
    }

    async function submit(){
      if(passwordsMatch()){
        const res = await fetch("http://127.0.0.1:3000/api/signup",{
            method:"POST",
            body: JSON.stringify({
                  username: username,
                  password: password
              })
          })
        console.log(res)
      }
    }
</script>

<section class="bg-gray-50 min-h-screen flex items-center justify-center">
  <div class="relative h-full w-full bg-white">
    <!-- Background -->
    <div class="absolute inset-0 bg-[linear-gradient(to_right,#4f4f4f2e_1px,transparent_1px),linear-gradient(to_bottom,#4f4f4f2e_1px,transparent_1px)] bg-[size:14px_24px] [mask-image:radial-gradient(ellipse_80%_50%_at_50%_0%,#000_70%,transparent_110%)]"></div>
    
    <div class="relative z-10 flex flex-col items-center mt-12 px-6 mx-auto md:h-screen lg:py-0">
      <div class="w-full bg-white rounded-lg shadow-md md:mt-0 sm:max-w-md xl:p-0 border border-gray-500">
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
              class="w-full text-black border border-gray-500 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
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
  </div>
</section>


