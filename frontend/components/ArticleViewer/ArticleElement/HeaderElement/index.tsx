import { Heading } from '@chakra-ui/react'
import parse from 'html-react-parser'

const HeaderElement = (articleDic: any) => {
  const level: any = `h${articleDic.data.level}`
  const size = (level: string) => {
    switch (level) {
      case 'h1':
        return '2xl'
      case 'h2':
        return 'xl'
      case 'h3':
        return 'lg'
      case 'h4':
        return 'md'
      case 'h5':
        return 'sm'
      case 'h6':
        return 'xs'
    }
  }
  return (
    <Heading as={level} size={size(level)}>
      {parse(articleDic.data.text)}
    </Heading>
  )
}
export default HeaderElement
