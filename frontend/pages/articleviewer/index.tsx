import { Grid, GridItem } from '@chakra-ui/react'
import ArticleColumn from '../../components/ArticleViewer/ArticleColumn'
import SideColumn from '../../components/ArticleViewer/SideColumn'
import Header from '../../components/Header'

const articleJson = {
  time: 1679211491245,
  blocks: [
    { id: 'DkMnOlbgD6', type: 'header', data: { text: 'Header', level: 1 } },
    { id: 'zCtAI2J2i5', type: 'code', data: { code: 'code\nCode()' } },
    { id: 'zSRWSRpDMZ', type: 'delimiter', data: {} },
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
  return (
    <>
      <Header></Header>
      <Grid templateColumns='repeat(4, 1fr)'>
        <GridItem colSpan={3}>
          <ArticleColumn articleJson={articleJson} />
        </GridItem>
        <GridItem colSpan={1}>
          <SideColumn />
        </GridItem>
      </Grid>
    </>
  )
}

export default ArticleViewer
