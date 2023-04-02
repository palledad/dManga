import Header from '@editorjs/header'
import CheckList from '@editorjs/checklist'
import Code from '@editorjs/code'
import Delimiter from '@editorjs/delimiter'
import Embed from '@editorjs/embed'
import InlineCode from '@editorjs/inline-code'
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
  checklist: CheckList,
  code: Code,
  delimiter: Delimiter,
  embed: Embed,
  inlineCode: InlineCode,
  linkTool: LinkTool,
  list: List,
  quote: Quote,
  paragraph: Paragraph,
}

export default Tools
