const evaluateToxicity = async (msg) => {
    try{
        const model = await Toxicity.load(0.8)
        const predictions = await model.classify(msg.text)
        let matches = predictions.filter( (p) => p.results[0].match === true );
        msg.isToxic = matches.length > 0 ? true : false
        msg.textFinal = msg.isToxic ? String.fromCodePoint(0x1F6AB).repeat(3) : msg.text
    }catch(error){
        console.log('--- ERROR evaluateToxicity() (Skipping) ---')
        msg.isToxic = null
        msg.textFinal = msg.text
        console.error(error)
    }
}