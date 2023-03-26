import parse from 'html-react-parser'
import React from 'react'
import { Code } from '@chakra-ui/react'

const CodeBlock = (articleDic: any) => {
  return <Code display='block'>{parse(articleDic.data.code)}</Code>
}
export default CodeBlock
