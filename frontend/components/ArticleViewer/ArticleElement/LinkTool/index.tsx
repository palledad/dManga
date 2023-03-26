import { ExternalLinkIcon } from '@chakra-ui/icons'
import { Link } from '@chakra-ui/react'

const LinkTool = (articleDic: any) => {
  return (
    <Link href={articleDic.data.link} isExternal>
      {articleDic.data.link}
      <ExternalLinkIcon mx='2px' />
    </Link>
  )
}
export default LinkTool
