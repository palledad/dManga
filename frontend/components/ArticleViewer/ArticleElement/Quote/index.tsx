import { Box } from '@chakra-ui/react'

const Quote = (articleDic: any) => {
  return (
    <Box my='1em' pl='2em' borderLeft='2px' lineHeight='1.8'>
      {articleDic.data.text}
    </Box>
  )
}
export default Quote
