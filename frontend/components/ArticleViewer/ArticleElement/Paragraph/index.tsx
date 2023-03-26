import parse from 'html-react-parser'
import React from 'react'
import { Text } from '@chakra-ui/react'

const Paragraph = (articleDic: any) => {
  return <Text>{parse(articleDic.data.text)}</Text>
}

export default Paragraph
