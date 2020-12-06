// This code is mostly based on Martin "arp242" Tournoij's work. See https://www.arp242.net/wasm-cli.html for more.
var hist = [], hist_index = 0;
(function() {
	window.set_output = (output) => {
		// Write stdout to terminal.
		let outputBuf = '';
		const decoder = new TextDecoder("utf-8");
		global.fs.writeSync = (fd, buf) => {
			outputBuf += decoder.decode(buf);
			const nl = outputBuf.lastIndexOf("\n");
			if (nl != -1) {
				output.innerText += outputBuf.substr(0, nl + 1);
				window.scrollTo(0, document.body.scrollHeight);
				outputBuf = outputBuf.substr(nl + 1);
			}
			return buf.length;
		};
	};
	window.readline = function(prompt, output, input, cb) {
		input.addEventListener('keydown', (e) => {
			// ^L
			if (e.ctrlKey && e.keyCode === 76) {
				e.preventDefault();
				output.innerText = '';
			}
			// ^P, arrow up
			else if ((e.ctrlKey && e.keyCode === 80) || e.keyCode === 38) {
				e.preventDefault();
				input.value = hist[hist.length - hist_index - 1] || '';
				if (hist_index < hist.length - 1)
					hist_index++;
			}
			// Arrow down; no ^N as it seems that can't be overridden: https://stackoverflow.com/q/38838302
			else if (e.keyCode === 40) {
				e.preventDefault();
				input.value = hist[hist.length - hist_index] || '';
				if (hist_index > 0)
					hist_index--;
			}
			// Enter
			else if (e.keyCode === 13) {
                e.preventDefault();
                output.innerText += prompt;
                output.innerText += input.value + "\n";
                hist.push(input.value);
                var cmd = input.value + '\n';
                input.value = '';
                if (cmd.length === 0)
                    return;
                cb(cmd);
            }
		});
		input.focus();
	};
}());
(function() {
    if (!WebAssembly.instantiateStreaming) {
        WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer();
            return await WebAssembly.instantiate(source, importObject);
        };
    }
    set_output(window.output);
    fetch('main.wasm').then(response => response.arrayBuffer()).then(function(bin) {
        const go = new Go();
        WebAssembly.instantiate(bin, go.importObject).then((result) => {
            go.run(result.instance);
            var prompt = 'λ > ';
            readline(prompt, window.output, window.input, repl);
            if (location.hash !== '') {
                var e = location.hash.substr(1);
                var cmd = atob(e) + "\n";
                output.innerText += prompt + cmd;
                console.log("d", e, cmd);
                repl(cmd);
            }
        });
    });
})();
(function() {
    var elem = document.querySelector('.permalink')
    elem.addEventListener('click', (e) => {
        e.preventDefault()
        var cmd = input.value
        if (cmd === "") {
            cmd = hist[hist.length-1]
        }
        if (cmd === "" || cmd == undefined ) {
            return alert('No term to link to.')
        }
        location.hash=btoa(cmd)
        alert("Permalink for "+cmd+":\n"+location)
        console.log("e", cmd, btoa(cmd))
        return input.focus()
    });
})();



