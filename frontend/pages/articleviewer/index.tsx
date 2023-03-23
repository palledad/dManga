import { Grid, GridItem } from '@chakra-ui/react'
import Header from '../../components/Header'
import ArticleColumn from '../../components/ArticleViewer/ArticleColumn'
import SideColumn from '../../components/ArticleViewer/SideColumn'
import { useEffect, useState } from 'react'
import parse from 'html-react-parser'

const articleJson = {
  time: 1679211491245,
  blocks: [
    { id: 'DkMnOlbgD6', type: 'header', data: { text: 'Header', level: 1 } },
    {
      id: 'ksw0x1Q4xu',
      type: 'checklist',
      data: {
        items: [
          { text: 'checklist', checked: false },
          { text: 'checked checklist', checked: true },
        ],
      },
    },
    { id: 'zCtAI2J2i5', type: 'code', data: { code: 'code\nCode()' } },
    {
      id: 'xGL_zO_NGu',
      type: 'linkTool',
      data: { link: 'https://developer.mozilla.org/ja/', meta: {} },
    },
    { id: 'JmYhyKv5wm', type: 'list', data: { style: 'ordered', items: ['list1', 'list2'] } },
    {
      id: 'VbVvlQmRoq',
      type: 'quote',
      data: { text: 'quote', caption: 'caption', alignment: 'left' },
    },
    { id: 'FxyUayqKcD', type: 'paragraph', data: { text: 'あああああああああ' } },
    { id: 'CiXQRCob8Z', type: 'paragraph', data: { text: 'aaaa' } },
  ],
  version: '2.26.5',
}

const ArticleViewer = () => {
  const [content, setContent] = useState([])

  const elementGenerator = (articleDic: any, key: number) => {
    const type = articleDic.type

    //switch文にする
    if (type === 'paragraph') {
      return <p key={key}>{parse(articleDic.data.text)}</p>
    } else {
      return <div key={key}>未対応</div>
    }
  }
  useEffect(() => {
    articleJson.blocks.map((dic, i) => elementGenerator(dic, i))
  }, [])
  return (
    <>
      <Header></Header>
      <Grid templateColumns='repeat(4, 1fr)'>
        <GridItem colSpan={3}>
          <div>{articleJson.blocks.map((dic, i) => elementGenerator(dic, i))}</div>
          <ArticleColumn />
        </GridItem>
        <GridItem colSpan={1}>
          <SideColumn />
        </GridItem>
      </Grid>
    </>
  )
}

export default ArticleViewer
