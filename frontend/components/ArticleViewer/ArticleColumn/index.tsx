import { Box } from '@chakra-ui/react'
import Segment from '../../Segment'
import Paragraph from '../ArticleElement/Paragraph'
import CodeBlock from '../ArticleElement/CodeBlock'
import HeaderElement from '../ArticleElement/HeaderElement'
import Delimiter from '../ArticleElement/Delimiter'
import LinkTool from '../ArticleElement/LinkTool'
import List from '../ArticleElement/ListElement'
import Quote from '../ArticleElement/Quote'
import React from 'react'

type Props = {
  articleJson: { [index: string]: any }
}
const ArticleColumn: React.FC<Props> = ({ articleJson }) => {
  const elementGenerator = (articleDic: any) => {
    switch (articleDic.type) {
      case 'paragraph':
        return Paragraph(articleDic)
      case 'code':
        return CodeBlock(articleDic)
      case 'header':
        return HeaderElement(articleDic)
      case 'delimiter':
        return Delimiter()
      case 'linktool':
        return LinkTool(articleDic)
      case 'list':
        return List(articleDic)
      case 'quote':
        return Quote(articleDic)
      default:
        return <p>未対応</p>
    }
  }

  return (
    <>
      <Box height={'20em'} background={'gray.100'}>
        サムネイル画像
      </Box>
      <Segment>
        <Box>記事タイトル</Box>
      </Segment>
      <Segment>
        <Box>タグ一覧</Box>
      </Segment>
      <Segment>
        <Box>
          {JSON.parse(articleJson.content).blocks
            ? JSON.parse(articleJson.content).blocks.map((dic: any) => elementGenerator(dic))
            : ''}
        </Box>
      </Segment>
    </>
  )
}
export default ArticleColumn
