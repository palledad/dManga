import Header from '@editorjs/header'
import Code from '@editorjs/code'
import Delimiter from '@editorjs/delimiter'
import Embed from '@editorjs/embed'
import LinkTool from '@editorjs/link'
import List from '@editorjs/list'
import Quote from '@editorjs/quote'
import Paragraph from '@editorjs/paragraph'

const Tools = {
  header: {
    class: Header,
    config: {
      placeholder: 'Enter a header',
      levels: [1, 2, 3],
      defaultLevel: 1,
    },
  },
  code: Code,
  delimiter: Delimiter,
  embed: Embed,
  linkTool: LinkTool,
  list: List,
  quote: Quote,
  paragraph: Paragraph,
}

export default Tools
