<script>
    import Counter from "../Counter.svelte";
    import {onMount} from "svelte";
    import {goto} from "$app/navigation";

    let username = ""

    onMount(() => {
        window.Twitch.ext.onAuthorized(async (auth) => {
                try {
                    console.log(auth)

                    console.log(window.Twitch.ext.viewer.id)

                    const user = await (await fetch(`https://api.twitch.tv/helix/users?id=${window.Twitch.ext.viewer.id}`, {
                            headers: {
                                "client-id": "r3q3lkzfdu6u0b9x4n82fmc0v3hyvf",
                                Authorization: `Extension ${auth.helixToken}`
                            }
                        }
                    )).json()
                    console.log(user)
                    console.log(user.data[0].display_name)
                    username = user.data[0].display_name

                    throw "adf"
                }
                catch (e) {
                    await goto("/error")
                }
            }
        )

    })
</script>

<h1>
    {username}
</h1>

<Counter/>