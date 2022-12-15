```
 __  __     ______   __         __   __    
/\_\_\_\   /\  == \ /\ \       /\ "-.\ \   
\/_/\_\/_  \ \  _-/ \ \ \____  \ \ \-.  \  
  /\_\/\_\  \ \_\    \ \_____\  \ \_\\"\_\ 
  \/_/\/_/   \/_/     \/_____/   \/_/ \/_/      
```

`xpln` is a command line interface to explain blocks of code using @OpenAI's GPT-3 model built using Go and the @spf13's Cobra library.

## Build

### Locally

```bash
git clone git@github.com:5amCurfew/xpln.git
```

Ensure you have Go (version 1.19) installed and an @OpenAI API token (refer to beta.openai.com) and the variable `OPENAI_API_KEY` set to your token in your shell

Build the binary the current directory using `go build` (can then be found using `./xpln`)

### Global

TODO: publish to Homebrew

## Usage

```bash
./xpln <PATH_TO_FILE> [OPTIONAL] --lines start-end
```

`xpln` requires the relative file path as the only argument. Optional flags include `--lines` (`-l`) that takes a string with the format `start-end` (inclusive, e.g. `5-12` will read from line 5 up-to-and-including line 12)

If no lines are provided `xpln` will read the entire file

See `xpln --help` for more details

## Example

Explain the `example.js` file in the repository
```
./xpln example.js
 ✓  Xpln'd
┌─ xpln ───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
| ┌─ Code Block ────────────────────────────────┐ ┌─ Explained ──────────────────────────────────────────────────────────┐     |
| | const evaluateToxicity = async (msg) => ... | | 1. Load the model                                                    |     |
| |     try{                                    | | 2. Classify the message                                              |     |
| |         const model = await Toxicity.loa... | | 3. Filter out toxic messages                                         |     |
| |         const predictions = await model.... | | 4. Set the isToxic flag to true or false                             |     |
| |                                             | | based on the number of toxic messages found in step 3                |     |
| |         let matches = predictions.filter... | | 5. Set the textFinal flag to a string of three                       |     |
| |                                             | | repeated characters if isToxic is true, otherwise set it to msg.text |     |
| |         msg.isToxic = matches.length > 0... | |                                                                      |     |
| |         msg.textFinal = msg.isToxic ? St... | └──────────────────────────────────────────────────────────────────────┘     |
| |     }catch(error){                          |                                                                              |
| |         console.log('--- ERROR evaluateT... |                                                                              |
| |         msg.isToxic = null                  |                                                                              |
| |         msg.textFinal = msg.text            |                                                                              |
| |         console.error(error)                |                                                                              |
| |     }                                       |                                                                              |
| | }                                           |                                                                              |
| |                                             |                                                                              |
| |                                             |                                                                              |
| └─────────────────────────────────────────────┘                                                                              |
|                                                                                                                              |
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘ 
```

## License

[MIT](https://choosealicense.com/licenses/mit/)