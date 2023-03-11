import Header from '../../components/Header'
import Segment from '../../components/Segment'
import { Box } from '@chakra-ui/react'
import dynamic from 'next/dynamic'
import { useState } from 'react'
const CustomEditor = dynamic(() => import('../../components/EditorJS/Editor'), {
  ssr: false,
})

const EditArticle = () => {
  const [data, setData] = useState()

  return (
    <>
      <Header></Header>
      <Segment>
        <Box>記事作成画面</Box>
      </Segment>
      <Segment>
        <Box bg='gray.500' w='100%' p={4}>
          <CustomEditor data={data} onChange={setData} holder='editorjs-container' />
        </Box>
      </Segment>
    </>
  )
}
export default EditArticle
