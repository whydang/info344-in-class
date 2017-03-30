# INFO 344 Spring 2017 In-Class Coding

Fork this repo into your own account. We will use this repo for all of the coding we do in-class. I will add starter code to the `master` branch of [my original repo](https://github.com/info344-s17/info344-in-class) as we go along, and you will pull those commits into your forked copy. At the end of class, you can commit the code you wrote and push it up to GitHub. I will also commit any code I write during class to the [completed branch](https://github.com/info344-s17/info344-in-class/tree/completed) in my original repo.

To pull updates, you'll need to add a new remote to your cloned fork. If you are using your own laptop during class, you should only have to do this once after cloning your fork. If you are using a lab machine, you'll have to do this each time you clone your forked repo to the lab machine. Run this command from within your repo directory:

```bash
git remote add upstream https://github.com/info344-s17/info344-in-class.git
```

When I ask you to "pull updates from the upstream master," execute this command from within your repo directory:

```bash
git pull upstream master
```

This will pull all commits I've made to the `master` branch of my original repo since you last pulled them.
