
## Server Demo/Example

### Stack:
* golang
* react (starting with create-react-app)
* mongodb
* tailwind
* webpack


##  godotenv
### Why we use it:
From the original (rails-dotenv) Library:

    Storing configuration in the environment is one of the tenets of a twelve-factor app. Anything that is likely to change between deployment environments–such as resource handles for databases or credentials for external services–should be extracted from the code into environment variables.

    But it is not always practical to set environment variables on development machines or continuous integration servers where multiple projects are run. Dotenv load variables from a .env file into ENV when the environment is bootstrapped.


<table>
<thead>
<tr>
<th>Hierarchy Priority</th>
<th>Filename</th>
<th>Environment</th>
<th>Should I <code>.gitignore</code>it?</th>
<th>Notes</th>
</tr>
</thead>
<tbody>
<tr>
<td>1st (highest)</td>
<td><code>.env.development.local</code></td>
<td>Development</td>
<td>Yes!</td>
<td>Local overrides of environment-specific settings.</td>
</tr>
<tr>
<td>1st</td>
<td><code>.env.test.local</code></td>
<td>Test</td>
<td>Yes!</td>
<td>Local overrides of environment-specific settings.</td>
</tr>
<tr>
<td>1st</td>
<td><code>.env.production.local</code></td>
<td>Production</td>
<td>Yes!</td>
<td>Local overrides of environment-specific settings.</td>
</tr>
<tr>
<td>2nd</td>
<td><code>.env.local</code></td>
<td>Wherever the file is</td>
<td>Definitely.</td>
<td>Local overrides. This file is loaded for all environments <em>except</em> <code>test</code>.</td>
</tr>
<tr>
<td>3rd</td>
<td><code>.env.development</code></td>
<td>Development</td>
<td>No.</td>
<td>Shared environment-specific settings</td>
</tr>
<tr>
<td>3rd</td>
<td><code>.env.test</code></td>
<td>Test</td>
<td>No.</td>
<td>Shared environment-specific settings</td>
</tr>
<tr>
<td>3rd</td>
<td><code>.env.production</code></td>
<td>Production</td>
<td>No.</td>
<td>Shared environment-specific settings</td>
</tr>
<tr>
<td>Last</td>
<td><code>.env</code></td>
<td>All Environments</td>
<td>Depends (See <a href="#should-i-commit-my-env-file">below</a>)</td>
<td>The Original®</td>
</tr>
</tbody>
</table>



ReactJS	Svelte
Popularity		
Github Stars	217k	74.8k
Github Contributors	1645	658
Github # Used By	14M	250k
StackOverflow Questions	472k	5.6k
NPM Weekly Downloads	7M	250k
		
Jetbrains Survey		
Developers Using	57.00%	7.00%
		
		
Release Date	2013	2016
Type	Library	Compiled
Size	42KB	1.7KB

