package main

// Given a string and a style array render HTML
// pretty much like a rich text editor.
// For example: 'Hello,world', [[0, 2, 'i'], [4, 9, 'b'], [7, 10, 'u']]
// Output: '<i>Hel</i>l<b>o,w<u>orl</u></b><u>d</u>

type Tag struct {
	Open   int
	Close  int
	TagStr string
}

func StyleParse(s string, arr []Tag) string {

}

/*
function format(str, arr) {
  const mapp = new Array(str.length);
  const strSplits = str.split('');

  arr.forEach(([o, c, t]) => {
    mapp[o] = [t, 'OPEN'];
    mapp[c] = [t, 'CLOSE'];
  });

  let ret = '';

  let stack = [];
  strSplits.forEach((c, i) => {
    const [openTag, isOpen] = mapp[i] ? mapp[i] : [];
    const [closeTag, isClose] = mapp[i] ? mapp[i] : [];

    let openStr = '';
    if (isOpen === 'OPEN') {
      openStr = '<' + openTag + '>'

      stack.push(openTag);
    }

    const currentlyOpen = stack[stack.length - 1];

    let closeStr = ''
    if (isClose === 'CLOSE') {
      if (currentlyOpen === closeTag) {
        stack.pop();
        closeStr = '</' + closeTag + '>';
      } else {
        const last = stack.pop();

        closeStr += '</' + last + '>';
        closeStr += '</' + closeTag + '>';
        closeStr += '<' + last + '>';

        stack.pop();
        stack.push(last);
      }
    }

    ret += openStr + c + closeStr
  });

  return ret;
}
*/
